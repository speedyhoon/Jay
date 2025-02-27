package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"reflect"
	"runtime"
	"strings"
)

// TODO add support for pointers with all types.
// TODO add support for enums with restricted sizes like: buf[1] = WriteEnum44(e.Enum1, e.Enum2)

// makeMarshal ...
func (s *structTyp) makeMarshal(b *bytes.Buffer, importJ *bool) {
	s.varsMarshal()
	varLengths := s.generateLenVarLine()
	makeSize := s.generateMakeSizes(s.calcSize())
	s.isReturnedInline()

	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)
	s.makeWriteBools(buf, &byteIndex, importJ)
	s.writeSingles(buf, &byteIndex, importJ)

	for _, f := range s.fixedLen {
		at := utl.UtoA(byteIndex)
		bufWriteLine(buf, f.marshalLine(&byteIndex, at, "", importJ, ""))
	}

	at, end := s.defineTrackingVars2(buf, byteIndex)
	for i, f := range s.stringSlice {
		if i >= 1 {
			at, end = f.track2(buf, i, len(s.stringSlice), end)
		}
		bufWriteLine(buf, f.marshalLine(&byteIndex, at, end, importJ, string(f.marshal.qtyVar)))
	}

	if at != vAt && end != vEnd {
		at, end = s.defineTrackingVars(buf, byteIndex)
	}
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		bufWriteLine(buf, f.marshalLine(&byteIndex, at, end, importJ, string(f.marshal.qtyVar)))
	}

	if len(buf.Bytes()) == 0 {
		return
	}

	var pointer string
	if s.option.PointerMarshalFunc {
		pointer = "*"
	}

	if s.returnInline {
		bufWriteF(b,
			"\nfunc (%s %s%s) MarshalJ() []byte {\n\treturn %s\n}\n",
			s.receiver,
			pointer,
			s.name,
			buf.String(),
		)
		return
	}

	bufWriteF(b,
		"\nfunc (%[1]s %[2]s%[3]s) MarshalJ() (%[4]s []byte) {\n\t%[5]s\n\t%[4]s = make([]byte, %[6]s)\n\t%[7]s%[8]s\treturn\n}\n",
		s.receiver,
		pointer,
		s.name,
		s.bufferName,
		varLengths,
		makeSize,
		s.generateSizeLine(),
		buf.String(),
	)

	return
}

func (s *structTyp) generateSizeLine() string {
	qty := len(s.variableLen)
	if qty == 0 {
		return ""
	}
	assignments, values := make([]string, qty), make([]string, qty)
	for i := 0; i < qty; i++ {
		assignments[i] = fmt.Sprintf("%s[%d]", s.bufferName, i)
		values[i] = fmt.Sprintf("byte(%s)", s.variableLen[i].marshal.qtyVar)
	}
	return fmt.Sprintln(strings.Join(assignments, ", "), "=", strings.Join(values, ", "))
}

func (s *structTyp) isReturnedInline() {
	s.returnInline = !(len(s.variableLen) >= 1 || len(s.fixedLen) >= 1 || len(s.stringSlice) >= 1) ||
		len(s.fixedLen) == 1 && s.fixedLen[0].isArray() && s.fixedLen[0].typ == tBytes
}

func (f *field) marshalLine(byteIndex *uint, at, end string, importJ *bool, lenVar string) string {
	fun, template := f.marshalFuncTemplate()
	totalSize := f.typeFuncSize()
	if template == tNoTemplate || template > tByteConv {
		// Unknown type, not supported yet.
		return ""
	}

	*byteIndex += totalSize
	thisField := f.Name()
	if end == "" {
		end = utl.UtoA(*byteIndex)
	}

	if f.isDef && fun != copyKeyword && f.isNotArrayOrSlice() {
		f.structTyp.imports.add(f.pkgReq)
		thisField = printFunc(f.typ, thisField)
	}

	switch template {
	case tFunc:
		if f.isArray() {
			thisField += "[:]"
		}

		return fmt.Sprintf("%s(%s, %s)", fun, f.sliceExpr(at, end), thisField)
	case tFuncOpt:
		return fmt.Sprintf("if %s != 0 {\n%s(%s, %s)\n}", lenVar, fun, f.sliceExpr(at, end), thisField)
	case tFuncLength:
		return fmt.Sprintf("%s(%s, %s, %s)", fun, f.sliceExpr(at, end), thisField, lenVar)
	case tByteAssign:
		return fmt.Sprintf("%s[:]", thisField)
	default:
		lg.Printf("template %d unhandled", template)
		return ""
	}
}

func (f *field) isNotArrayOrSlice() bool {
	return f.arraySize == typeNotArrayOrSlice
}
func (f *field) isArray() bool {
	return f.arraySize >= typeArray
}
func (f *field) isSlice() bool {
	return f.arraySize <= typeSlice
}

func printFunc(fun string, params ...string) (code string) {
	if fun == "" {
		return strings.Join(params, ", ")
	}
	code = fmt.Sprintf("%s(%s)", fun, strings.Join(params, ", "))
	return
}

func (f *field) marshalFuncTemplate() (funcName string, template uint8) {
	var fun any
	switch f.typ {
	case tBools:
		fun, template = jay.WriteBools, tFuncLength
	case tByte:
		if f.isDef {
			return tByte, tByteConv
		}
		return "", tByteAssign
	case tBytes:
		if f.isArray() {
			if f.isFirst && f.isLast {
				return "", tByteAssign
			}
			return copyKeyword, tFunc
		}

		if f.Required {
			return copyKeyword, tFunc
		}
		return copyKeyword, tFuncOpt
	case tFloat32:
		fun, template = jay.WriteFloat32, tFunc
	case tFloat32s:
		fun, template = jay.WriteFloat32s, tFuncLength
	case tFloat64:
		fun, template = jay.WriteFloat64, tFunc
	case tFloat64s:
		fun, template = jay.WriteFloat64s, tFuncLength
	case tInt16:
		fun, template = jay.WriteInt16, tFunc
	case tInt16s:
		fun, template = jay.WriteInt16s, tFuncLength
	case tInt32:
		fun, template = jay.WriteInt32, tFunc
	case tInt32s:
		fun, template = jay.WriteInt32s, tFunc
	case tInt64:
		fun, template = jay.WriteInt64, tFunc
	case tInt64s:
		fun, template = jay.WriteInt64s, tFunc
	case tInt8:
		return tByte, tByteConv
	case tInt8s:
		fun, template = jay.WriteInt8s, tFunc
	case tInt:
		if f.structTyp.option.FixedIntSize {
			if f.structTyp.option.Is32bit {
				fun, template = jay.WriteIntX32, tFunc
			}
			fun, template = jay.WriteIntX64, tFunc
		} else {
			fun, template = jay.WriteIntVariable, tFuncLength
		}
	case tInts:
		if f.structTyp.option.Is32bit {
			fun, template = jay.WriteIntsX32, tFunc
		} else {
			fun, template = jay.WriteIntsX64, tFunc
		}
	case tString:
		return copyKeyword, tFunc
	case tStrings:
		fun, template = f.sizeOfPick(jay.WriteStrings8, jay.WriteStrings16), tFunc
	case tTime:
		if f.tagOptions.TimeNano {
			fun, template = jay.WriteTimeNano, tFunc
		} else {
			fun, template = jay.WriteTime, tFunc
		}
	case tTimeDurations:
		fun, template = jay.WriteDurations, tFunc
	case tUint16:
		fun, template = jay.WriteUint16, tFunc
	case tUint16s:
		fun, template = jay.WriteUint16s, tFuncLength
	case tUint32:
		fun, template = jay.WriteUint32, tFunc
	case tUint32s:
		fun, template = jay.WriteUint32s, tFunc
	case tUint64:
		fun, template = jay.WriteUint64, tFunc
	case tUint64s:
		fun, template = jay.WriteUint64s, tFunc
	case tUint:
		if f.structTyp.option.FixedUintSize {
			if f.structTyp.option.Is32bit {
				fun, template = jay.WriteUintX32, tFunc
			}
			fun, template = jay.WriteUintX64, tFunc
		} else {
			fun, template = jay.WriteUintVariable, tFuncLength
		}
	case tUints:
		if f.structTyp.option.Is32bit {
			fun, template = jay.WriteUintsX32, tFunc
		} else {
			fun, template = jay.WriteUintsX64, tFunc
		}

	default:
		lg.Printf("no function set for type %s yet in typeFuncs()", f.typ)
		return
	}
	return nameOf(fun, f.structTyp.isImportJ), template
}

func nameOf(f any, importJ *bool) string {
	if s, ok := f.(string); ok {
		return s
	}

	if importJ != nil {
		*importJ = true
	}

	s := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), "/")
	return s[len(s)-1]
}

func (f *field) pickSizeFunc(small, large any) string {
	return nameOf(f.sizeOfPick(small, large), f.structTyp.isImportJ)
}

func (f *field) sliceExpr(at, end string) string {
	if f.typ == tStrings && f.isFirst && f.isLast {
		return f.structTyp.bufferName
	}

	if at == "0" {
		at = ""
	}

	if f.isFixedLen {
		if f.isFirst && f.isLast {
			return f.structTyp.bufferName
		}

		// `at == ""` is needed when structType contains variableLen types
		// then `at` can't be absent because their sizes are placed before.
		if f.isFirst && at == "" {
			return fmt.Sprintf("%s[:%s]", f.structTyp.bufferName, end)
		}
	}

	if f.isLast {
		return fmt.Sprintf("%s[%s:]", f.structTyp.bufferName, at)
	}
	return fmt.Sprintf("%s[%s:%s]", f.structTyp.bufferName, at, end)
}

func (f *field) sliceExpr2(at, end string, byteIndex uint) string {
	if f.typ == tStrings && f.isFirst && f.isLast {
		return f.structTyp.bufferName
	}

	if at == "0" {
		at = ""
	}

	if f.isFixedLen {
		if f.isFirst && f.isLast {
			return f.structTyp.bufferName
		}

		// `at == ""` is needed when structType contains variableLen types
		// then `at` can't be absent because their sizes are placed before.
		if f.isFirst && at == "" {
			return fmt.Sprintf("%s[:%s]", f.structTyp.bufferName, end)
		}
	}

	if f.isLast {
		if at == "" && byteIndex == 0 {
			return f.structTyp.bufferName
		}
		if at == "" && byteIndex != 0 {
			return fmt.Sprintf("%s[%d:]", f.structTyp.bufferName, byteIndex)
		}
		return fmt.Sprintf("%s[%s:]", f.structTyp.bufferName, at)
	}
	return fmt.Sprintf("%s[%s:%s]", f.structTyp.bufferName, at, end)
}

func (f *field) sliceExpr3(c *varCtx) string {
	if c.atValue == "" && c.endValue == "" || f.isFirst && f.isLast && *f.indexStart == 0 {
		return f.structTyp.bufferName
	}

	if f.isLast || !f.isFixedLen && f.isStrings() {
		return fmt.Sprintf("%s[%s:]", f.structTyp.bufferName, c.atValue)
	}

	if f.isFixedLen {
		return fmt.Sprintf("%s[%s:%d]", f.structTyp.bufferName, omitZero(*f.indexStart), *f.indexEnd)
	}

	return fmt.Sprintf("%s[%s:%s]", f.structTyp.bufferName, c.atValue, c.endValue)
}

// Template definitions.
const (
	tNoTemplate uint8 = iota

	// tFunc call a function with the type, `func(b[at:end], type)`.
	tFunc

	// tFuncOpt wraps tFunc with an if statement. `if l0 != 0 { tFunc... }`
	tFuncOpt

	// tFuncLength calls a function with a length parameter, `func(b[at:end], type, int)`.
	tFuncLength

	// tFuncPtr calls a function with the field used as a pointer in the last parameter.
	tFuncPtr

	// tFuncPtrCheck wraps tFuncPtr with an if statement to check for any returned errors.
	tFuncPtrCheck

	tIfPtrCheck

	// tFuncPtrCheckAt expands tFuncPtrCheck by adding a function parameter to update a pointer for the byte index.
	tFuncPtrCheckAt

	// tByteAssign byte assignment always populated. No function calls required, `b[0] = byte`.
	tByteAssign

	// tByteConv converts that type to a byte & assigns, `b[0] = byte(int8)`.
	tByteConv
)
