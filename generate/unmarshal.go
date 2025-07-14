package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"strconv"
	"strings"
)

// moveReadStringsAbove is the threshold to move all `s.stringSlice` above `s.bool`, `s.single` and `s.fixedLen`.
// This reduces processing in generated unmarshall functions when errors occur by placing the buffer length
// checks occur before all the assignments.
const moveReadStringsAbove = 2

func (s *structTyp) putFixedLenBefore() bool {
	first := s.firstVarLenField()
	return len(s.stringSlice) < moveReadStringsAbove && first != nil && *first.indexStart <= 8
}

type varCtx struct {
	buf                       *bytes.Buffer
	atValue, endValue         string
	isAtDefined, isEndDefined bool
	isOkDefined               bool
}

func (c *varCtx) defineAtOk() string {
	if c.isAtDefined && c.isOkDefined {
		return fmt.Sprintf("%s, %s =", vAt, vOk)
	}
	c.isAtDefined, c.isOkDefined = true, true
	return fmt.Sprintf("%s, %s :=", vAt, vOk)
}

func (c *varCtx) defineAtLine() string {
	if !c.isAtDefined {
		return fmt.Sprintf("var %s uint\n\t", vAt)
	}
	return ""
}

// Names of variables used in the generated code.
const vAt, vEnd, vOk = "at", "end", "ok"

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer) {
	buf, fixedBuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	c := varCtx{buf: buf}
	s.varsUnmarshal()
	lengthChecks := s.makeLengthChecks()
	s.makeReadBools(fixedBuf)
	s.readSingles(fixedBuf)

	for _, f := range s.fixedLen {
		bufWriteLine(fixedBuf, f.unmarshalLine(&c))
	}

	isBefore := s.putFixedLenBefore()
	if isBefore {
		buf.Write(fixedBuf.Bytes())
	}

	for _, f := range s.stringSlice {
		bufWriteLine(buf, f.unmarshalLine(&c))
	}

	if !isBefore {
		// Place all fixed length types after stringSlice so the function can exit early
		// if encountering any unmarshalling errors.
		// `c.byteIndex` needs to be calculated before `s.stringSlice` which is why the execution order looks odd.
		buf.Write(fixedBuf.Bytes())
	}

	for _, f := range s.variableLen {
		bufWriteLine(buf, f.unmarshalLine(&c))
	}

	if len(buf.Bytes()) == 0 {
		return
	}

	if !s.returnInlineUnmarshal {
		bufWriteLine(buf, "return nil")
	}

	bufWriteF(b,
		"\nfunc (%s *%s) UnmarshalJ(%s []byte) error {\n\t%s\n%s}\n",
		s.receiver,
		s.name,
		s.bufferName,
		lengthChecks,
		buf.String(),
	)
}

func (s *structTyp) makeLengthChecks() string {
	qty := len(s.variableLen)
	if qty == 0 {
		return fmt.Sprintf("if len(%s) %s %d {\n\t\treturn %s\n\t}", s.bufferName, s.sizeCompSymbol(), s.qtyBytesRequired, s.ImportErr())
	}

	assignments, values := make([]string, 0, qty), make([]string, 0, qty)
	sizeGroups := varSize{}
	for i, f := range s.variableLen {
		if f.isLast && f.typ == tBools {
			sizeGroups.add(f, printFunc(f.pickSizeFunc(jay.SizeBools8, jay.SizeBools), string(f.unmarshal.qtyVar)))
			continue
		}

		assignments = append(assignments, string(f.unmarshal.sizeVar)) // Append the size variable name without any multiplier.
		if f.typ == tBools {
			values = append(values, printFunc(f.pickSizeFunc(jay.SizeBools8, jay.SizeBools), string(f.unmarshal.qtyVar)))
		} else {
			values = append(values, fmt.Sprintf("%s(%s)", intKeyword, f.qtyBytes()))
		}
		sizeGroups.add(f, assignments[i])
	}

	var assignLine string
	if len(assignments) >= 1 && len(values) >= 1 {
		assignLine = fmt.Sprintf(
			"%s := %s",
			strings.Join(assignments, ", "),
			strings.Join(values, ", "),
		)
	}

	if assignLine == "" {
		return fmt.Sprintf(
			`%[1]s := len(%[2]s)
	if %[1]s < %[3]d || %[1]s %[4]s %[3]d+%[5]s {
		return %[6]s
	}`,
			s.lengthName,
			s.bufferName,
			s.qtyBytesRequired,
			s.sizeCompSymbol(),
			sizeGroups.group(),
			s.ImportErr(),
		)
	}

	return fmt.Sprintf(
		`%[1]s := len(%[2]s)
	if %[1]s < %[3]d {
		return %[4]s
	}
	%[5]s
	if %[1]s %[6]s %[3]d+%[7]s {
		return %[4]s
	}`,
		s.lengthName,
		s.bufferName,
		s.qtyBytesRequired,
		s.ImportErr(),
		assignLine,
		s.sizeCompSymbol(),
		sizeGroups.group(),
	)
}

// sizeCompSymbol returns the less than symbol `<` when structTyp contains any []string types.
// Otherwise, a not equal to symbol `!=` is returned.
func (s *structTyp) sizeCompSymbol() string {
	if len(s.stringSlice) >= 1 {
		return "<"
	}
	return "!="
}

func (s *structTyp) generateMakeSizes(totalSize uint) string {
	sizeChecks := varSize{}
	for _, f := range append(s.stringSlice, s.variableLen...) {
		if f.isLast && f.typ == tStrings {
			if f.isArray() {
				sizeChecks.add(f, printFunc(nameOf(jay.SizeStringsArray, s.isImportJ), f.Field(""), strconv.Itoa(f.arraySize)))
			} else {
				sizeChecks.add(f, printFunc(nameOf(jay.SizeStrings8, s.isImportJ), f.Name()))
			}
		} else {
			if f.typ == tBools {
				sizeChecks.add(f, printFunc(nameOf(jay.SizeBools, s.isImportJ), string(f.marshal.qtyVar)))
			} else {
				sizeChecks.add(f, string(f.marshal.qtyVar))
			}
		}
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

// trackingVars defines `at` and `end` if needed before a value is unmarshalled or updates their values.
func (c *varCtx) trackingVars(f *field) {
	if f.isFirst && f.isLast && *f.indexStart == 0 {
		return
	}

	if f.isLast {
		if c.endValue != "" {
			c.atValue = c.endValue
			return
		}
		if !c.isAtDefined {
			c.atValue = omitZero(*f.indexStart)
		}
		return
	}

	if f.isFixedLen {
		c.atValue, c.endValue = omitZero(*f.indexStart), omitZero(*f.indexEnd)
		return
	}

	if f.isStrings() {
		if f.isFirst && !f.isArray() {
			c.isAtDefined, c.atValue = true, vAt
			return
		}

		if !c.isAtDefined {
			if len(f.structTyp.stringSlice) >= 2 || len(f.structTyp.variableLen) >= 1 {
				c.isAtDefined = true
				if f.indexStart != nil && *f.indexStart != 0 {
					bufWriteLineF(c.buf, "%s := %d", vAt, *f.indexStart)
				} else {
					bufWriteLineF(c.buf, "var %s uint", vAt)
				}
				c.atValue = vAt
				if !c.isEndDefined {
					c.endValue = ""
				}
			}
		}
		return
	}

	if f.isVarLen() {
		if c.isAtDefined && c.isEndDefined {
			c.atEndLineInc(f.unmarshal.sizeVar.String(f.elmSize))
			return
		}

		if !c.isAtDefined {
			c.atEndLineSet(*f.indexStart, f.unmarshal.sizeVar.String(f.elmSize))
			return
		}

		if !c.isEndDefined {
			bufWriteLineF(c.buf, "%s := %s+%s", vEnd, vAt, f.unmarshal.sizeVar.String(f.elmSize))
			c.isEndDefined = true
			c.endValue = vEnd
		}
	}
}

func omitZero(u uint) string {
	if u == 0 {
		return ""
	}
	return utl.UtoA(u)
}

func (c *varCtx) atEndLineInc(inc any) {
	atEndLineInc(c.buf, inc)
}

// atEndLineInc increments existing `at` and `end` variables.
func atEndLineInc(buf *bytes.Buffer, inc any) {
	bufWriteLineF(buf, "%s, %s = %[2]s, %[2]s+%v", vAt, vEnd, inc)
}

func (c *varCtx) atEndLineSet(byteIndex uint, lenVar string) {
	c.isAtDefined, c.isEndDefined = true, true
	atEndLineSet(c.buf, byteIndex, lenVar)
	c.atValue, c.endValue = vAt, vEnd
}

func atEndLineSet(buf *bytes.Buffer, byteIndex uint, lenVar string) {
	bufWriteLineF(buf, "%s, %s := %d, %[3]d+%v", vAt, vEnd, byteIndex, lenVar)
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
		return fmt.Sprintf("if %s != 0 {\n%s = %s\n}", f.unmarshal.qtyVar, f.Name(), f.sliceExpr3(ctx))

	case tFuncArrayTypes:
		return fmt.Sprintf("%s%s(%s, %s, %d)", f.structTyp.returnInlineUnmarshal, fun, f.sliceExpr3(ctx), f.Field(fun), f.arraySize)

	case tFuncLength:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(f.convertTo(), printFunc(fun, f.sliceExpr3(ctx), string(f.unmarshal.qtyVar))))

	case tFuncPtr:
		return fmt.Sprintf("%s%s(%s, &%s, %s)", f.structTyp.returnInlineUnmarshal, fun, f.sliceExpr3(ctx), f.Name(), f.qtyBytes())

	case tFuncPtrCheck:
		return fmt.Sprintf(
			"%s, %s := %s(%s, &%s, %s, %d)\n\tif !%[2]s {\n\t\treturn %[8]s\n\t}",
			vAt, vOk, fun, f.sliceExpr3(ctx), f.Name(), f.qtyBytes(), *f.indexStart, f.structTyp.ImportErr(),
		)

	case tFuncCheck:
		return fmt.Sprintf(
			"if !%s(%s, %s, %d, &%s) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr3(ctx), f.Field(fun), f.arraySize, vAt, f.structTyp.ImportErr(),
		)

	case tFuncCheck2:
		return fmt.Sprintf(
			"if !%s(%s, %s, %d, &%s) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr3(ctx), f.Field(fun), f.arraySize, vAt, f.structTyp.ImportErr(),
		)

	case tFuncPtrCheckAt:
		return fmt.Sprintf(
			"if !%s(%s, &%s, %s, &%s) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr3(ctx), f.Name(), f.qtyBytes(), vAt, f.structTyp.ImportErr(),
		)

	case tFuncPtrCheckAtOk:
		return fmt.Sprintf(
			"if !%s(%s, &%s, %s) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr3(ctx), f.Name(), f.qtyBytes(), f.structTyp.ImportErr(),
		)

	case tIfPtrCheck:
		return fmt.Sprintf(
			"if !%s(%s, &%s) {\n\t\treturn %s\n\t}",
			fun, f.sliceExpr3(ctx), f.Name(), f.structTyp.ImportErr(),
		)

	case tByteConv:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(fun, f.sliceExpr3(ctx)))

	case tByteAssign:
		return fmt.Sprintf("%s = %s", f.Name(), printFunc(f.convertTo(), f.sliceExpr3(ctx)))

	default:
		lg.Panicln("unhandled template")
		return ""
	}
}

func (f *field) qtyBytes() string {
	switch l := len(f.qtyIndex); l {
	case 0:
		// Used for array types.
		return f.structTyp.bufferName
	case 1:
		return f.indexExp(&f.qtyIndex[0])
	default:
		return f.sliceExp(&f.qtyIndex[0], &f.qtyIndex[l-1])
	}
}

func (f *field) indexExp(start *uint) string {
	if start != nil {
		return fmt.Sprintf("%s[%d]", f.structTyp.bufferName, *start)
	}
	return fmt.Sprintf("%s", f.structTyp.bufferName)
}

func (f *field) sliceExp(start, end *uint) string {
	var s, e string
	if start != nil && *start != 0 {
		s = utl.UtoA(*start)
	}
	if end != nil && *end != 0 {
		e = utl.UtoA(*end)
	}
	if s == "" && e == "" {
		return fmt.Sprintf("%s", f.structTyp.bufferName)
	}
	return fmt.Sprintf("%s[%s:%s]", f.structTyp.bufferName, s, e)
}

func (f *field) qtySlice() string {
	l := len(f.qtyIndex)
	if l == 0 {
		panic("f.qtyIndex not populated")
	}

	end := f.qtyIndex[l-1] + 1
	return f.sliceExp(&f.qtyIndex[0], &end)
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
		if f.aliasType != "" {
			f.structTyp.imports.add(f.pkgReq)
			return fmt.Sprintf("[%d]%s", f.arraySize, f.aliasType)
		} else {
			return fmt.Sprintf("[%d]%s", f.arraySize, f.arrayType)
		}
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
func (f *field) unmarshalFunc() (funcName string, template uint8, canReturnInline canReturnInlined) {
	var c interface{}
	switch f.typ {
	case tBools:
		c, template = jay.ReadBools8, tFuncLength
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
	case tBytes:
		if f.isArray() {
			c, template = "", tByteAssign
			return
		}
		if f.Required {
			c, template = "", tByteAssign
		} else {
			c, template = copyKeyword, tFuncOpt
		}
	case tComplex64:
		c, template = jay.ReadComplex64, tFunc
	case tComplex64s:
		c, template = jay.ReadComplex64s, tFuncLength
	case tComplex128:
		c, template = jay.ReadComplex128, tFunc
	case tComplex128s:
		c, template = jay.ReadComplex128s, tFuncLength
	case tFloat32:
		c, template = jay.ReadFloat32, tFunc
	case tFloat32s:
		c, template = jay.ReadFloat32s, tFuncLength
	case tFloat64:
		c, template = jay.ReadFloat64, tFunc
	case tFloat64s:
		c, template = jay.ReadFloat64s, tFuncLength
	case tInt16:
		c, template = jay.ReadInt16, tFunc
	case tInt16s:
		c, template = jay.ReadInt16s, tFuncLength
	case tInt32:
		c, template = jay.ReadInt32, tFunc
	case tInt32s:
		c, template = jay.ReadInt32s, tFuncLength
	case tInt64:
		c, template = jay.ReadInt64, tFunc
	case tInt64s:
		c, template = jay.ReadInt64s, tFuncLength
	case tInt8s:
		c, template = jay.ReadInt8s, tFuncLength
	case tInt:
		if f.structTyp.option.VariableIntSize {
			c, template = jay.ReadInt, tFunc
		} else {
			if f.structTyp.option.Is32bit {
				c, template = jay.ReadIntX32, tFunc
			} else {
				c, template = jay.ReadIntX64, tFunc
			}
		}
	case tInts:
		if f.structTyp.option.Is32bit {
			c, template = jay.ReadIntsX32, tFuncLength
		} else {
			c, template = jay.ReadIntsX64, tFuncLength
		}
	case tStrings:
		if f.isArray() {
			if f.isLast {
				c, template, canReturnInline = jay.ReadStringsArrayErr, tFuncArrayTypes, canReturnInlined(f.isLast)
			} else if f.isFirst {
				c, template, canReturnInline = jay.ReadStringsArrayAtOk, tFuncCheck, canReturnInlined(f.isLast)
			} else {
				c, template, canReturnInline = jay.ReadStringsArrayAtOk, tFuncCheck2, canReturnInlined(f.isLast)
			}
		} else if f.isLast {
			canReturnInline = canReturnInlined(f.structTyp.putFixedLenBefore())
			if canReturnInline {
				c, template, canReturnInline = f.sizeOfPick(jay.ReadStrings8Err, jay.ReadStrings8Err), tFuncPtr, canReturnInlined(f.structTyp.putFixedLenBefore())
			} else {
				c, template = f.sizeOfPick(jay.ReadStrings8Ok, jay.ReadStrings8Ok), tFuncPtrCheckAtOk
			}
		} else if f.isFirst {
			c, template = f.sizeOfPick(jay.ReadStrings8nbXt, jay.ReadStrings8nbXt), tFuncPtrCheck
		} else {
			c, template = f.sizeOfPick(jay.ReadStrings8nbX, jay.ReadStrings8nbX), tFuncPtrCheckAt
		}
	case tTime:
		if f.tagOptions.TimeNano {
			c, template = jay.ReadTimeNano, tFunc
		} else {
			c, template = jay.ReadTime, tFunc
		}
	case tTimes:
		c, template = jay.ReadTimes, tFuncLength
	case tTimeDurations:
		c, template = jay.ReadDurations, tFuncLength
	case tUint16:
		c, template = jay.ReadUint16, tFunc
	case tUint16s:
		c, template = jay.ReadUint16s, tFuncLength
	case tUint32:
		c, template = jay.ReadUint32, tFunc
	case tUint32s:
		c, template = jay.ReadUint32s, tFuncLength
	case tUint64:
		c, template = jay.ReadUint64, tFunc
	case tUint64s:
		c, template = jay.ReadUint64s, tFuncLength
	case tUint:
		if f.structTyp.option.VariableUintSize {
			c, template = jay.ReadUintVariable, tFunc
		} else {
			if f.structTyp.option.Is32bit {
				c, template = jay.ReadUintX32, tFunc
			} else {
				c, template = jay.ReadUintX64, tFunc
			}
		}
	case tUints:
		if f.structTyp.option.Is32bit {
			c, template = jay.ReadUintsX32, tFuncLength
		} else {
			c, template = jay.ReadUintsX64, tFuncLength
		}

	default:
		lg.Printf("no function set for type %s yet in unmarshalFunc()", f.typ)
	}

	return nameOf(c, f.structTyp.isImportJ), template, canReturnInline
}
