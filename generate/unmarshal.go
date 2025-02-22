package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"strings"
)

// moveReadStringsAbove is the threshold to move all `s.stringSlice` above `s.bool`, `s.single` and `s.fixedLen`.
// This reduces processing in generated unmarshall functions when errors occur by placing the buffer length
// checks occur before all the assignments.
const moveReadStringsAbove = 2

func (s *structTyp) ReturnInline() bool {
	return len(s.fixedLen) == 0 && len(s.single) == 0 && len(s.variableLen) == 0 && len(s.bool) == 0 && len(s.stringSlice) >= 1
}

func (s *structTyp) placeStringsAfter(byteIndex uint) bool {
	return len(s.stringSlice) < moveReadStringsAbove && byteIndex <= 2
}

type varCtx struct {
	byteIndex                 uint
	size                      uint
	buf                       *bytes.Buffer
	atValue, endValue         string
	isAtDefined, isEndDefined bool
}

const vAt, vEnd, vOk = "at", "end", "ok"

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer) {
	buf, fixedBuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	c := varCtx{byteIndex: uint(len(s.variableLen)), buf: buf}

	s.makeReadBools(fixedBuf, &c.byteIndex)
	s.readSingles(fixedBuf, &c.byteIndex)

	for _, f := range s.fixedLen {
		fixedBuf.WriteString(f.unmarshalLine(&c))
		fixedBuf.WriteString("\n")
	}

	isAfter := s.placeStringsAfter(c.byteIndex)
	if isAfter {
		buf.Write(fixedBuf.Bytes())
	}

	for _, f := range s.stringSlice {
		buf.WriteString(f.unmarshalLine(&c))
		buf.WriteString("\n")
	}

	if !isAfter {
		// Place all fixed length types after stringSlice so the function can exit early
		// if encountering any unmarshalling errors.
		// `c.byteIndex` needs to be calculated before `s.stringSlice` which is why the execution order looks odd.
		buf.Write(fixedBuf.Bytes())
	}

	for _, f := range s.variableLen {
		buf.WriteString(f.unmarshalLine(&c))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	// Prevent panic: runtime error: index out of range
	var lengthCheck string
	if len(s.variableLen) == 0 && len(s.stringSlice) == 0 {
		lengthCheck = fmt.Sprintf("if len(%s) != %d {\nreturn %s\n}", s.bufferName, c.byteIndex, exportedErr)
	} else if len(s.variableLen) == 0 && len(s.stringSlice) >= 1 {
		lengthCheck = fmt.Sprintf("if len(%s) < %d {\nreturn %s\n}", s.bufferName, c.byteIndex, exportedErr)
	} else {
		lengthCheck = fmt.Sprintf("%[1]s := len(%[2]s)\nif %[1]s < %[3]d {\nreturn %[4]s\n}", s.lengthName, s.bufferName, c.byteIndex, exportedErr)
	}
	variableLengthCheck := s.generateCheckSizes(c.byteIndex)

	if s.ReturnInline() {
		bufWriteF(b,
			"func (%s *%s) UnmarshalJ(%s []byte) error {\n%s\n%s\n}\n",
			s.receiver,
			s.name,
			s.bufferName,
			lengthCheck,
			code,
		)
		return
	}

	if s.returnInlineUnmarshal {
		bufWriteF(b,
			"func (%s *%s) UnmarshalJ(%s []byte) error {\n%s\n%s%s}\n",
			s.receiver,
			s.name,
			s.bufferName,
			lengthCheck,
			variableLengthCheck,
			code,
		)
		return
	}

	bufWriteF(b,
		"func (%s *%s) UnmarshalJ(%s []byte) error {\n%s\n%s%sreturn nil\n}\n",
		s.receiver,
		s.name,
		s.bufferName,
		lengthCheck,
		variableLengthCheck,
		code,
	)
}

func (s *structTyp) generateCheckSizes(totalSize uint) string {
	qty := len(s.variableLen)
	if qty == 0 {
		return ""
	}

	assignments, values := make([]string, qty), make([]string, qty)
	sizeChecks := make(varSize, 5) // 1,2,4,8 and 0 (bool)
	for i, f := range s.variableLen {
		assignments[i] = f.lenVar
		values[i] = fmt.Sprintf("int(%s[%d])", s.bufferName, i)
		sizeChecks.add(f, assignments[i])
	}

	return fmt.Sprintf(
		"%s := %s\nif %s != %d+%s {\nreturn %s\n}\n",
		strings.Join(assignments, ", "),
		strings.Join(values, ", "),
		s.lengthName,
		totalSize,
		sizeChecks.group(),
		exportedErr,
	)
}

func (s *structTyp) generateMakeSizes(totalSize uint) string {
	sizeChecks := make(varSize, 5) // 1,2,4,8 and 0 (bool)
	for _, f := range append(s.stringSlice, s.variableLen...) {
		sizeChecks.add(f, f.lenVar)
	}

	grouped := sizeChecks.group()
	if totalSize == 0 {
		return grouped
	}
	if grouped == "" {
		return utl.UtoA(totalSize)
	}

	return fmt.Sprintf("%d+%s", totalSize, grouped)
}

func (f *field) isStrings() bool {
	return f.fieldList == &f.structTyp.stringSlice
}
func (f *field) isVarLen() bool {
	return f.fieldList == &f.structTyp.variableLen
}

func (c *varCtx) deferAdd(f *field, size uint) {
	if f.isFixedLen || f.typ == tStrings {
		c.byteIndex += size
	}
}

// trackingVars defines `at` and `end` if needed before a value is unmarshalled or updates their values.
func (c *varCtx) trackingVars(f *field) {
	size := f.typeFuncSize()
	defer c.deferAdd(f, size)

	if f.isFirst && f.isLast {
		return
	}

	if f.isLast {
		if c.endValue != "" {
			c.atValue = c.endValue
			return
		}
		if !c.isAtDefined {
			c.atValue = utl.UtoA(c.byteIndex)
		}
		return
	}

	if f.isFixedLen {
		if c.byteIndex >= 1 {
			c.atValue = utl.UtoA(c.byteIndex)
		}
		c.endValue = utl.UtoA(c.byteIndex + size)
		return
	}

	if f.isStrings() {
		if f.isFirst {
			c.isAtDefined, c.atValue = true, vAt
			return
		}

		if !c.isAtDefined {
			if len(f.structTyp.stringSlice) >= 2 || len(f.structTyp.variableLen) >= 1 {
				c.isAtDefined = true
				c.buf.WriteString(fmt.Sprintf("%s := %d\n", vAt, c.byteIndex))
				c.atValue = vAt
			}
		}
		return
	}

	if f.isVarLen() {
		if !c.isAtDefined && len(f.structTyp.variableLen) >= 3 {
			c.isAtDefined, c.isEndDefined = true, true
			c.buf.WriteString(fmt.Sprintf("%s, %s := %d, %[3]d+%s\n", vAt, vEnd, c.byteIndex, f.lenVar))
			c.atValue, c.endValue = vAt, vEnd
		}

		if !c.isEndDefined && len(f.structTyp.variableLen) >= 3 {
			c.isEndDefined = true
			c.buf.WriteString(fmt.Sprintf("%s := %s+%s\n", vEnd, vAt, f.lenVar))
			c.endValue = vEnd
		} else {
			c.endValue = fmt.Sprintf("%s+%s", c.atValue, f.lenVar)
		}
	}
}

func (f *field) unmarshalLine(ctx *varCtx) string {
	var fun string
	var template uint8
	fun, template, f.structTyp.returnInlineUnmarshal = f.unmarshalFunc()

	ctx.trackingVars(f)

	switch template {
	case tFunc:
		if f.isDef && f.isNotArrayOrSlice() {
			return fmt.Sprintf("%s = %s", f.Name(), printFunc(f.aliasType, printFunc(fun, f.sliceExpr3(ctx))))
		}
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(fun, f.sliceExpr3(ctx)))

	case tFuncOpt:
		return fmt.Sprintf("if %s != 0 {\n%s = %s\n}", f.lenVar, f.Name(), f.sliceExpr3(ctx))

	case tFuncLength:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(fun, f.sliceExpr3(ctx), f.lenVar))

	case tFuncPtr:
		return fmt.Sprintf("%s%s(%s, &%s)", f.structTyp.returnInlineUnmarshal, fun, f.sliceExpr3(ctx), f.Name())

	case tFuncPtrCheck:
		return fmt.Sprintf(
			"%s, %s := %s(%s, &%s)\n\tif !%[2]s {\n\t\treturn %[6]s\n\t}",
			vAt, vOk, fun, f.structTyp.bufferName, f.Name(), exportedErr,
		)

	case tFuncPtrCheckAt:
		return fmt.Sprintf(
			"if !%s(%s, &%s, &%s) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr3(ctx), f.Name(), vAt, exportedErr,
		)

	case tIfPtrCheck:
		return fmt.Sprintf(
			"if !%s(%s, &%s) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr3(ctx), f.Name(), exportedErr,
		)

	case tByteConv:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(fun, f.sliceExpr3(ctx)))

	case tByteAssign:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(f.convertTo(), f.sliceExpr3(ctx)))
	}

	lg.Println("unhandled template")
	return ""
}

type canReturnInlined bool

func (c canReturnInlined) String() string {
	if c {
		return "return "
	}
	return ""
}

func (f *field) convertTo() string {
	if f.isArray() {
		return fmt.Sprintf("[%d]%s", f.arraySize, f.arrayType)
	}
	return ""
}

func (f *field) typeConvert() string {
	if f.pkgReq != "" {
		f.structTyp.imports.add(f.pkgReq)
		return pkgSelName(f.pkgReq, f.aliasType)
	}
	return f.aliasType
}

// unmarshalFunc returns the function name to handle unmarshalling.
// `size` is the quantity of bytes required to represent the type.
func (f field) unmarshalFunc() (funcName string, template uint8, canReturnInline canReturnInlined) {
	var c interface{}
	switch f.typ {
	case tByte:
		if f.isDef {
			return f.typeConvert(), tByteConv, false
		}
		return "", tByteAssign, false
	case tInt8, tString:
		if f.isDef {
			return f.typeConvert(), tByteConv, false
		}
		return f.typ, tByteConv, false
	case tInt:
		if f.structTyp.option.FixedIntSize {
			if f.structTyp.option.Is32bit {
				c, template = jay.ReadIntX32, tFunc
			}
			c, template = jay.ReadIntX64, tFunc
			break
		}
		//c = jay.ReadIntVariable
		c, template = jay.ReadInt, tFunc
	case tInt16:
		c, template = jay.ReadInt16, tFunc
	case tInt32:
		c, template = jay.ReadInt32, tFunc
	case tInt64:
		c, template = jay.ReadInt64, tFunc
	case tFloat32:
		c, template = jay.ReadFloat32, tFunc
	case tFloat64:
		c, template = jay.ReadFloat64, tFunc
	case tUint:
		if f.structTyp.option.FixedUintSize {
			if f.structTyp.option.Is32bit {
				c, template = jay.ReadUintX32, tFunc
			}
			c, template = jay.ReadUintX64, tFunc
			break
		}
		c, template = jay.ReadUintVariable, tFunc
	case tUint16:
		c, template = jay.ReadUint16, tFunc
	case tUint32:
		c, template = jay.ReadUint32, tFunc
	case tUint64:
		c, template = jay.ReadUint64, tFunc
	case tTime:
		if f.tagOptions.TimeNano {
			c, template = jay.ReadTimeNano, tFunc
		} else {
			c, template = jay.ReadTime, tFunc
		}
	case tInt8S:
		c, template = jay.ReadInt8s, tFuncLength
	case tBoolS:
		c, template = jay.ReadBools, tFuncLength

	case tByteS:
		if f.isArray() {
			c, template = "", tByteAssign
			return
		}

		if f.Required {
			c, template = "", tByteAssign
		} else {
			c, template = copyKeyword, tFuncOpt
		}
	case tFloat32S:
		c, template = jay.ReadFloat32s, tFuncLength
	case tFloat64S:
		c, template = jay.ReadFloat64s, tFuncLength
	case tInt16S:
		c, template = jay.ReadInt16s, tFuncLength
	case tUintS:
		if f.structTyp.option.Is32bit {
			c, template = jay.ReadUintsX32, tFuncLength
		}
		c, template = jay.ReadUintsX64, tFuncLength
	case tUint16S:
		c, template = jay.ReadUint16s, tFuncLength
	case tIntS:
		if f.structTyp.option.Is32bit {
			c, template = jay.ReadIntsX32, tFuncLength
		}
		c, template = jay.ReadIntsX64, tFuncLength
	case tInt32S:
		c, template = jay.ReadInt32s, tFuncLength
	case tInt64S:
		c, template = jay.ReadInt64s, tFuncLength
	case tUint32S:
		c, template = jay.ReadUint32s, tFuncLength
	case tUint64S:
		c, template = jay.ReadUint64s, tFuncLength
	case tTimeDurations:
		c, template = jay.ReadDurations, tFuncLength
	case tStrings:
		if f.isLast {
			c, template, canReturnInline = f.sizeOfPick(jay.ReadStrings8Err, jay.ReadStrings16Err), tFuncPtr, canReturnInlined(f.isLast)
		} else if f.isFirst {
			c, template = f.sizeOfPick(jay.ReadStrings8n, jay.ReadStrings16n), tFuncPtrCheck
		} else {
			c, template = f.sizeOfPick(jay.ReadStrings8nb, jay.ReadStrings16nb), tFuncPtrCheckAt
		}

	default:
		lg.Printf("no function set for type %s yet in unmarshalFunc()", f.typ)
	}

	return nameOf(c, nil), template, canReturnInline
}
