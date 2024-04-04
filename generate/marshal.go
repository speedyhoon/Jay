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
func (s *structTyp) makeMarshal(b *bytes.Buffer, o Option, importJ *bool) {
	varLengths := lengths2(s.varLenFieldNames(o), s.receiver)
	makeSize := joinSizes(s.calcSize(o), s.variableLen, o, importJ)

	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)
	isReturnInlined := s.makeWriteBools(buf, &byteIndex, importJ)
	isReturnInlined = s.writeSingles(buf, &byteIndex, s.receiver, o, importJ) || isReturnInlined

	for _, f := range s.fixedLen {
		buf.WriteString(o.generateLine(s, f, &byteIndex, Utoa(byteIndex), "", importJ, ""))
		buf.WriteString("\n")
	}

	at, end := s.defineTrackingVars(buf, byteIndex)
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(o.generateLine(s, f, &byteIndex, at, end, importJ, lenVariable(i)))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	var pointer string
	if o.PointerMarshalFunc {
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

func (o Option) generateLine(s *structTyp, f field, byteIndex *uint, at, end string, importJ *bool, lenVar string) string {
	fun, template, totalSize := f.MarshalFuncTemplate(o, importJ)
	if fun == "" {
		// Unknown type, not supported yet.
		return ""
	}

	*byteIndex += totalSize
	thisField := pkgSelName(s.receiver, f.name)
	if end == "" {
		end = Utoa(*byteIndex)
	}

	if f.typ != f.aliasType && f.aliasType != "" && fun != copyKeyword {
		s.imports.add(f.pkgReq)
		thisField = printFunc(f.typ, thisField)
	}

	switch template {
	case tFunc:
		return fmt.Sprintf("%s(%s, %s)", fun, sliceExpr(s, f, at, end), thisField)
	case tFuncOpt:
		return fmt.Sprintf("if %s != 0 {\n%s(%s, %s)\n}", lenVar, fun, sliceExpr(s, f, at, end), thisField)
	case tFuncLength:
		return fmt.Sprintf("%s(%s, %s, %s)", fun, sliceExpr(s, f, at, end), thisField, lenVar)
	default:
		lg.Printf("template %d unhandled", template)
		return ""
	}
}

func (f *field) isArrayOrSlice() bool {
	return f.arraySize != 0
}
func (f *field) isArray() bool {
	return f.arraySize >= 1
}
func (f *field) isSlice() bool {
	return f.arraySize <= -1
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

func (f field) MarshalFuncTemplate(o Option, importJ *bool) (funcName string, template uint8, totalSize uint) {
	switch f.typ {
	case "byte", "uint8":
		if f.typ != f.aliasType {
			return "byte", tByteConv, 1
		}
		return "", tByteAssign, 1
	case "int8":
		return "byte", tByteConv, 1
	case "string":
		return copyKeyword, tFunc, 0
	case "[]uint8", "[]byte":
		if f.Required {
			return copyKeyword, tFunc, 0
		}
		return copyKeyword, tFuncOpt, 0
	case "[15]byte", "[15]uint8":
		return copyKeyword, tFunc, uint(f.arraySize)
	}

	var fun any
	fun, template, totalSize = f.marshalFunc(o)
	return nameOf(fun, importJ), template, totalSize
}

func (f field) marshalFunc(o Option) (fun interface{}, template uint8, totalSize uint) {
	switch f.typ {
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				return jay.WriteIntArch32, tFunc, 4
			}
			return jay.WriteIntArch64, tFunc, 8
		}
		return jay.WriteIntVariable, tFuncLength, 1
	case "int16":
		return jay.WriteInt16, tFunc, 2
	case "int32", "rune":
		return jay.WriteInt32, tFunc, 4
	case "float32":
		return jay.WriteFloat32, tFunc, 4
	case "float64":
		return jay.WriteFloat64, tFunc, 8
	case "int64":
		return jay.WriteInt64, tFunc, 8
	case "time.Duration":
		return jay.WriteDuration, tFunc, 8
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				return jay.WriteUintArch32, tFunc, 4
			}
			return jay.WriteUintArch64, tFunc, 8
		}
		return jay.WriteUintVariable, tFuncLength, 1
	case "uint16":
		return jay.WriteUint16, tFunc, 2
	case "uint32":
		return jay.WriteUint32, tFunc, 4
	case "uint64":
		return jay.WriteUint64, tFunc, 8
	case "time.Time":
		if f.tagOptions.TimeNano {
			return jay.WriteTimeNano, tFunc, 8
		} else {
			return jay.WriteTime, tFunc, 8
		}
	case "[]int8":
		return jay.WriteInt8s, tFunc, 1
	case "[]bool":
		return jay.WriteBools, tFuncLength, 1

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

func sliceExpr(s *structTyp, f field, at, end string) string {
	if at == "0" {
		at = ""
	}

	if f.isFixedLen {
		if f.isFirst && f.isLast {
			return s.bufferName
		}
		if f.isFirst && at == "" { // `at == ""` is needed when structType contains variableLen types then `at` can't be absent because their sizes are placed before.
			return fmt.Sprintf("%s[:%s]", s.bufferName, end)
		}
	}

	if f.isLast {
		return fmt.Sprintf("%s[%s:]", s.bufferName, at)
	}
	return fmt.Sprintf("%s[%s:%s]", s.bufferName, at, end)
}

func sliceExprU(s *structTyp, f field, at, end uint) string {
	return sliceExpr(s, f, Utoa(at), Utoa(end))
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
