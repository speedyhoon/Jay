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
	fun, totalSize := o.typeFuncs(f, importJ)
	if fun == "" {
		// Unknown type, not supported yet.
		lg.Printf("no generateLine for type `%s` yet", f.typ)
		return ""
	}

	start := *byteIndex
	*byteIndex += totalSize
	thisField := pkgSelName(s.receiver, f.name)

	switch f.typ {
	case "string":
		return fmt.Sprintf("%s(%s, %s)", fun, sliceExpr(s, f, at, end), thisField)
	case "[]byte", "[]uint8":
		if f.isFirst && f.isLast {
			if f.Required {
				return fmt.Sprintf("%s(%s, %s)", fun, sliceExpr(s, f, at, end), thisField)
			}
			return fmt.Sprintf("if %s != 0 {\n%s(%s, %s)\n}", lenVar, fun, sliceExpr(s, f, at, end), thisField)
		} else {
			if f.Required {
				return fmt.Sprintf("%s(%s, %s)", fun, sliceExpr(s, f, at, end), thisField)
			}
			return fmt.Sprintf("if %s != 0 {%s(%s, %s)\n}", lenVar, fun, sliceExpr(s, f, at, end), thisField)
		}
	case "[]int8":
		return fmt.Sprintf("%s(%s, %s)", fun, sliceExpr(s, f, at, end), thisField)
	case "[]bool":
		return fmt.Sprintf("%s(%s, %s, %s)", fun, sliceExpr(s, f, at, end), thisField, lenVar)
	}

	if f.isArray() && fun == "copy" {
		thisField += "[:]"
	}

	if f.typ != f.aliasType {
		s.imports.add(f.pkgReq)
		thisField = printFunc(f.typ, thisField)
	}
	return printFunc(fun, sliceExpr(s, f, Utoa(start), Utoa(*byteIndex)), thisField)
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

/*func marshalU64(f *field) (fun string, sizeOf uint, ok bool) {
	if f.tagOptions.maxBytes != 0 {
		return fmt.Sprintf("WriteUint%dBytes", f.tagOptions.maxBytes*8), f.tagOptions.maxBytes, true
	}
	return "", 0, false
}*/

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

func (o Option) typeFuncs(fe field, importJ *bool) (fun string, totalSize uint) {
	var f interface{}
	switch fe.typ {
	case "byte", "uint8":
		if fe.typ != fe.aliasType {
			return "byte", 1
		}
		return "", 1
	case "int8":
		return "byte", 1
	case "string":
		return "copy", 0
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				f, totalSize = jay.WriteIntArch32, 4
			}
			f, totalSize = jay.WriteIntArch64, 8
			break
		}
		f, totalSize = jay.WriteIntVariable, 1
	case "int16":
		f, totalSize = jay.WriteInt16, 2
	case "int32", "rune":
		f, totalSize = jay.WriteInt32, 4
	case "float32":
		f, totalSize = jay.WriteFloat32, 4
	case "float64":
		f, totalSize = jay.WriteFloat64, 8
	case "int64":
		f, totalSize = jay.WriteInt64, 8
	case "time.Duration":
		f, totalSize = jay.WriteDuration, 8
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				f, totalSize = jay.WriteUintArch32, 4
			}
			f, totalSize = jay.WriteUintArch64, 8
			break
		}
		f, totalSize = jay.WriteUintVariable, 1
	case "uint16":
		f, totalSize = jay.WriteUint16, 2
	case "uint32":
		f, totalSize = jay.WriteUint32, 4
	case "uint64":
		f, totalSize = jay.WriteUint64, 8
	case "time.Time":
		if fe.tagOptions.TimeNano {
			f, totalSize = jay.WriteTimeNano, 8
		} else {
			f, totalSize = jay.WriteTime, 8
		}
	case "[]uint8", "[]byte":
		return "copy", 0
	case "[]int8":
		f, totalSize = jay.WriteInt8s, 1
	case "[]bool":
		f, totalSize = jay.WriteBools, 1

	case "[15]byte", "[15]uint8":
		return "copy", uint(fe.arraySize)

	default:
		log.Printf("no function set for type %s yet in typeFuncs()", fe.typ)
		return "", 0, 0
	}

	return nameOf(f, importJ), totalSize
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
		if f.isFirst {
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
