package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"strings"
)

func (s *structTyp) ReturnInline() bool {
	return len(s.fixedLen) == 0 && len(s.single) == 0 && len(s.variableLen) == 0 && len(s.bool) == 0 && len(s.stringSlice) >= 1
}

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer) {
	var byteIndex = uint(len(s.variableLen))
	fixedBuf := bytes.NewBuffer(nil)

	s.makeReadBools(fixedBuf, &byteIndex)
	s.readSingles(fixedBuf, &byteIndex)

	for _, f := range s.fixedLen {
		fixedBuf.WriteString(f.unmarshalLine(&byteIndex, utl.UtoA(byteIndex), "", ""))
		fixedBuf.WriteString("\n")
	}

	buf := bytes.NewBuffer(nil)
	var at, end string
	if len(s.stringSlice) >= 1 {
		at, end = s.defineTrackingVars(buf, byteIndex)
	}
	for i, f := range s.stringSlice {
		at, _ = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(f.unmarshalLine(&byteIndex, at, "", f.lenVar))
		buf.WriteString("\n")
	}

	// Place all fixed length types after stringSlice so the function can exit early
	// if encountering any unmarshalling errors.
	buf.Write(fixedBuf.Bytes())
	if len(s.stringSlice) == 0 {
		at, end = s.defineTrackingVars(buf, byteIndex)
	}
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(f.unmarshalLine(&byteIndex, at, end, f.lenVar))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	// Prevent panic: runtime error: index out of range
	var lengthCheck string
	if len(s.variableLen) == 0 && len(s.stringSlice) == 0 {
		lengthCheck = fmt.Sprintf("if len(%s) != %d {\nreturn %s\n}", s.bufferName, byteIndex, exportedErr)
	} else if len(s.variableLen) == 0 && len(s.stringSlice) >= 1 {
		lengthCheck = fmt.Sprintf("if len(%s) < %d {\nreturn %s\n}", s.bufferName, byteIndex, exportedErr)
	} else {
		lengthCheck = fmt.Sprintf("%[1]s := len(%[2]s)\nif %[1]s < %[3]d {\nreturn %[4]s\n}", s.lengthName, s.bufferName, byteIndex, exportedErr)
	}
	variableLengthCheck := s.generateCheckSizes(byteIndex)

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

func (f *field) unmarshalLine(byteIndex *uint, at, end, lenVar string) string {
	var fun string
	var template uint8
	fun, template, f.structTyp.returnInlineUnmarshal = f.unmarshalFuncs()
	totalSize := f.typeFuncSize()

	if f.isFixedLen || f.typ == tStrings {
		*byteIndex += totalSize
	}

	if end == "" && f.typ != tStrings {
		end = utl.UtoA(*byteIndex)
	}

	switch template {
	case tFunc:
		if f.isDef && f.isNotArrayOrSlice() {
			return fmt.Sprintf("%s = %s", f.Name(), printFunc(f.aliasType, printFunc(fun, f.sliceExpr(at, end))))
		}
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(fun, f.sliceExpr(at, end)))

	case tFuncOpt:
		return fmt.Sprintf("if %s != 0 {\n%s = %s\n}", lenVar, f.Name(), f.sliceExpr(at, end))

	case tFuncLength:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(fun, f.sliceExpr(at, end), lenVar))

	case tFuncPtr:
		return fmt.Sprintf("%s%s(%s, &%s)", f.structTyp.returnInlineUnmarshal, fun, f.sliceExpr(at, end), f.Name())

	case tFuncPtrCheck:
		return fmt.Sprintf(
			"at, ok := %s(%s, &%s)\n\tif !ok {\n\t\treturn %s\n\t}",
			fun, f.structTyp.bufferName, f.Name(), exportedErr,
		)

	case tFuncPtrCheckAt:
		return fmt.Sprintf(
			"if !%s(%s, &%s, &at) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr(at, end), f.Name(), exportedErr,
		)

	case tByteConv:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(fun, f.sliceExpr(at, end)))

	case tByteAssign:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(f.convertTo(), f.sliceExpr(at, end)))
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

// unmarshalFuncs returns the function name to handle unmarshalling.
// `size` is the quantity of bytes required to represent the type.
func (f field) unmarshalFuncs() (funcName string, template uint8, canReturnInline canReturnInlined) {
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
		lg.Printf("no function set for type %s yet in unmarshalFuncs()", f.typ)
	}

	return nameOf(c, nil), template, canReturnInline
}
