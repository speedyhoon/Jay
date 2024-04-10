package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// TODO add support for pointers with all types.
// TODO add support for enums with restricted sizes like: buf[1] = WriteEnum44(e.Enum1, e.Enum2)

// makeMarshal ...
func (s *structTyp) makeMarshal(b *bytes.Buffer, importJ *bool) {
	varLengths := lengths2(s.varLenFieldNames(), s.receiver)
	makeSize := joinSizes(s.calcSize(), s.variableLen, importJ)

	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)
	isReturnInlined := s.makeWriteBools(buf, &byteIndex, importJ)
	isReturnInlined = s.writeSingles(buf, &byteIndex, s.receiver, importJ) || isReturnInlined

	for _, f := range s.fixedLen {
		buf.WriteString(f.generateLine(&byteIndex, Utoa(byteIndex), "", importJ, ""))
		buf.WriteString("\n")
	}

	at, end := s.defineTrackingVars(buf, byteIndex)
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(f.generateLine(&byteIndex, at, end, importJ, lenVariable(i)))
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

	if isReturnInlined {
		bufWriteF(b,
			"func (%s %s%s) MarshalJ() []byte {\nreturn []byte{%s}\n}\n",
			s.receiver,
			pointer,
			s.name,
			code,
		)
		return
	}

	b.WriteString(fmt.Sprintf(
		"func (%[1]s %[8]s%[2]s) MarshalJ() (%[3]s []byte) {\n%[4]s\n%[3]s = make([]byte, %[5]s)\n%[6]s%[7]sreturn\n}\n",
		s.receiver,
		s.name,
		s.bufferName,
		varLengths,
		makeSize,
		s.generateSizeLine(),
		code,
		pointer,
	))

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

func (f field) generateLine(byteIndex *uint, at, end string, importJ *bool, lenVar string) string {
	fun, template := f.MarshalFuncTemplate(importJ)
	totalSize := f.typeFuncSize()
	if fun == "" {
		// Unknown type, not supported yet.
		return ""
	}

	*byteIndex += totalSize
	thisField := pkgSelName(f.structTyp.receiver, f.name)
	if end == "" {
		end = Utoa(*byteIndex)
	}

	if f.isAliasDef && fun != copyKeyword && f.arraySize == typeNotArrayOrSlice {
		f.structTyp.imports.add(f.pkgReq)
		thisField = printFunc(f.typ, thisField)
	}

	switch template {
	case tFunc:
		return fmt.Sprintf("%s(%s, %s)", fun, sliceExpr(f, at, end), thisField)
	case tFuncOpt:
		return fmt.Sprintf("if %s != 0 {\n%s(%s, %s)\n}", lenVar, fun, sliceExpr(f, at, end), thisField)
	case tFuncLength:
		return fmt.Sprintf("%s(%s, %s, %s)", fun, sliceExpr(f, at, end), thisField, lenVar)
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

// Utoa is equivalent to strconv.FormatUint(uint64(u), 10).
func Utoa(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
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
	case "byte", "uint8":
		if f.isAliasDef {
			return "byte", tByteConv
		}
		return "", tByteAssign
	case "int8":
		return "byte", tByteConv
	case "string":
		return copyKeyword, tFunc
	case "[]uint8", "[]byte":
		if f.Required {
			return copyKeyword, tFunc
		}
		return copyKeyword, tFuncOpt
	case "[15]byte", "[15]uint8":
		return copyKeyword, tFunc
	}

	var fun any
	fun, template = f.marshalFunc()
	return nameOf(fun, importJ), template
}

func (f field) marshalFunc() (fun interface{}, template uint8) {
	switch f.typ {
	case "int":
		if f.structTyp.option.FixedIntSize {
			if f.structTyp.option.Is32bit {
				return jay.WriteIntArch32, tFunc
			}
			return jay.WriteIntArch64, tFunc
		}
		return jay.WriteIntVariable, tFuncLength
	case "int16":
		return jay.WriteInt16, tFunc
	case "int32", "rune":
		return jay.WriteInt32, tFunc
	case "float32":
		return jay.WriteFloat32, tFunc
	case "float64":
		return jay.WriteFloat64, tFunc
	case "int64":
		return jay.WriteInt64, tFunc
	case "time.Duration":
		return jay.WriteDuration, tFunc
	case "uint":
		if f.structTyp.option.FixedUintSize {
			if f.structTyp.option.Is32bit {
				return jay.WriteUintArch32, tFunc
			}
			return jay.WriteUintArch64, tFunc
		}
		return jay.WriteUintVariable, tFuncLength
	case "uint16":
		return jay.WriteUint16, tFunc
	case "uint32":
		return jay.WriteUint32, tFunc
	case "uint64":
		return jay.WriteUint64, tFunc
	case "time.Time":
		if f.tagOptions.TimeNano {
			return jay.WriteTimeNano, tFunc
		} else {
			return jay.WriteTime, tFunc
		}
	case "[]int8":
		return jay.WriteInt8s, tFunc
	case "[]bool":
		return jay.WriteBools, tFuncLength
	case "[]float32":
		return jay.WriteFloat32s, tFuncLength

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

func sliceExpr(f field, at, end string) string {
	if at == "0" {
		at = ""
	}

	if f.isFixedLen {
		if f.isFirst && f.isLast {
			return f.structTyp.bufferName
		}
		if f.isFirst && at == "" { // `at == ""` is needed when structType contains variableLen types then `at` can't be absent because their sizes are placed before.
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

	copyKeyword = "copy"
)
