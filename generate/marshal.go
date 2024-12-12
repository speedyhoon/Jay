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
	varLengths := lengths2(s.varLenFieldNames(), s.receiver)
	makeSize := joinSizes(s.calcSize(), s.variableLen, importJ)
	s.isReturnedInline()

	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)
	s.makeWriteBools(buf, &byteIndex, importJ)
	s.writeSingles(buf, &byteIndex, s.receiver, importJ)

	for _, f := range s.fixedLen {
		at := utl.UtoA(byteIndex)
		buf.WriteString(f.marshalLine(&byteIndex, at, "", importJ, ""))
		buf.WriteString("\n")
	}

	at, end := s.defineTrackingVars(buf, byteIndex)
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(f.marshalLine(&byteIndex, at, end, importJ, lenVariable(i)))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	var pointer string
	if s.option.PointerMarshalFunc {
		pointer = "*"
	}

	if s.returnInline {
		bufWriteF(b,
			"func (%s %s%s) MarshalJ() []byte {\n\treturn %s\n}\n",
			s.receiver,
			pointer,
			s.name,
			code,
		)
		return
	}

	bufWriteF(b,
		"func (%[1]s %[2]s%[3]s) MarshalJ() (%[4]s []byte) {\n%[5]s\n%[4]s = make([]byte, %[6]s)\n%[7]s%[8]s\treturn\n}\n",
		s.receiver,
		pointer,
		s.name,
		s.bufferName,
		varLengths,
		makeSize,
		s.generateSizeLine(),
		code,
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
		values[i] = fmt.Sprintf("byte(%s)", lenVariable(i))
	}
	return fmt.Sprintln(strings.Join(assignments, ", "), " = ", strings.Join(values, ", "))
}

func (s *structTyp) isReturnedInline() {
	s.returnInline = !(len(s.variableLen) >= 1 || len(s.fixedLen) >= 1) ||
		len(s.fixedLen) == 1 && s.fixedLen[0].arraySize >= typeArray && s.fixedLen[0].typ == tByteS
}

func (f *field) marshalLine(byteIndex *uint, at, end string, importJ *bool, lenVar string) string {
	fun, template := f.MarshalFuncTemplate(importJ)
	totalSize := f.typeFuncSize()
	if template == 0 {
		// Unknown type, not supported yet.
		return ""
	}

	*byteIndex += totalSize
	thisField := pkgSelName(f.structTyp.receiver, f.name)
	if end == "" {
		end = utl.UtoA(*byteIndex)
	}

	if f.isDef && fun != copyKeyword && f.arraySize == typeNotArrayOrSlice {
		f.structTyp.imports.add(f.pkgReq)
		thisField = printFunc(f.typ, thisField)
	}

	switch template {
	case tFunc:
		if f.arraySize >= typeArray {
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

func (f *field) isArrayOrSlice() bool {
	return f.arraySize != typeNotArrayOrSlice
}
func (f *field) isArray() bool {
	return f.arraySize > typeNotArrayOrSlice
}
func (f *field) isSlice() bool {
	return f.arraySize == typeSlice
}

func printFunc(fun string, params ...string) (code string) {
	if fun == "" {
		return strings.Join(params, ", ")
	}
	code = fmt.Sprintf("%s(%s)", fun, strings.Join(params, ", "))
	return
}

func (f field) MarshalFuncTemplate(importJ *bool) (funcName string, template uint8) {
	switch f.typ {
	case tByte:
		if f.isDef {
			return tByte, tByteConv
		}
		return "", tByteAssign
	case tInt8:
		return tByte, tByteConv
	case tString:
		return copyKeyword, tFunc
	case tByteS:
		if f.arraySize >= typeArray {
			if f.isFirst && f.isLast {
				return "", tByteAssign
			}
			return copyKeyword, tFunc
		}

		if f.Required {
			return copyKeyword, tFunc
		}
		return copyKeyword, tFuncOpt
	}

	var fun any
	fun, template = f.marshalFunc()
	return nameOf(fun, importJ), template
}

func (f field) marshalFunc() (fun interface{}, template uint8) {
	switch f.typ {
	case tInt:
		if f.structTyp.option.FixedIntSize {
			if f.structTyp.option.Is32bit {
				return jay.WriteIntX32, tFunc
			}
			return jay.WriteIntX64, tFunc
		}
		return jay.WriteIntVariable, tFuncLength
	case tInt16:
		return jay.WriteInt16, tFunc
	case tInt32:
		return jay.WriteInt32, tFunc
	case tInt64:
		return jay.WriteInt64, tFunc
	case tFloat32:
		return jay.WriteFloat32, tFunc
	case tFloat64:
		return jay.WriteFloat64, tFunc
	case tUint:
		if f.structTyp.option.FixedUintSize {
			if f.structTyp.option.Is32bit {
				return jay.WriteUintX32, tFunc
			}
			return jay.WriteUintX64, tFunc
		}
		return jay.WriteUintVariable, tFuncLength
	case tUint16:
		return jay.WriteUint16, tFunc
	case tUint32:
		return jay.WriteUint32, tFunc
	case tUint64:
		return jay.WriteUint64, tFunc
	case tTime:
		if f.tagOptions.TimeNano {
			return jay.WriteTimeNano, tFunc
		} else {
			return jay.WriteTime, tFunc
		}
	case tInt8S:
		return jay.WriteInt8s, tFunc
	case tBoolS:
		return jay.WriteBools, tFuncLength
	case tFloat32S:
		return jay.WriteFloat32s, tFuncLength
	case tFloat64S:
		return jay.WriteFloat64s, tFuncLength
	case tInt16S:
		return jay.WriteInt16s, tFuncLength
	case tUintS:
		if f.structTyp.option.Is32bit {
			return jay.WriteUintsX32, tFunc
		}
		return jay.WriteUintsX64, tFunc
	case tUint16S:
		return jay.WriteUint16s, tFuncLength
	case tIntS:
		if f.structTyp.option.Is32bit {
			return jay.WriteIntsX32, tFunc
		}
		return jay.WriteIntsX64, tFunc
	case tInt32S:
		return jay.WriteInt32s, tFunc
	case tInt64S:
		return jay.WriteInt64s, tFunc
	case tUint32S:
		return jay.WriteUint32s, tFunc
	case tUint64S:
		return jay.WriteUint64s, tFunc

	default:
		lg.Printf("no function set for type %s yet in typeFuncs()", f.typ)
		return
	}
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

func (f *field) sliceExpr(at, end string) string {
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

// Template definitions.
const (
	// tFunc call a function with the type, `func(b[at:end], type)`.
	tFunc uint8 = iota + 1

	// tFuncOpt wraps tFunc with an if statement. `if l0 != 0 { tFunc... }`
	tFuncOpt

	// tFuncLength calls a function with a length parameter, `func(b[at:end], type, int)`.
	tFuncLength

	// tByteAssign byte assignment always populated. No function calls required, `b[0] = byte`.
	tByteAssign

	// tByteConv converts that type to a byte & assigns, `b[0] = byte(int8)`.
	tByteConv
)
