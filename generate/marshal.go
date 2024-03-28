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

	for i, f := range s.fixedLen {
		buf.WriteString(o.generateLine(s, f, &byteIndex, "", "", 0, i == 0, i == len(s.fixedLen)-1 && len(s.variableLen) == 0, importJ, ""))
		buf.WriteString("\n")
	}

	at, end := s.defineTrackingVars(buf, byteIndex)
	vLen := len(s.variableLen) - 1
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(o.generateLine(s, f, &byteIndex, at, end, uint(i), i == 0, i == vLen, importJ, lenVariable(i)))
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

func fieldNames(fields []field, receiver string) string {
	var s []string
	for i := range fields {
		s = append(s, pkgSelName(receiver, fields[i].name))
	}
	return strings.Join(s, ", ")
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

func (o Option) generateLine(s *structTyp, f field, byteIndex *uint, at, end string, index uint, isFirst, isLast bool, importJ *bool, lenVar string) string {
	fun, size, totalSize := o.typeFuncs(f, importJ)
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
		//if f.typ != f.aliasType {
		//	fun = f.aliasType
		//}
		if isFirst && isLast {
			return fmt.Sprintf("%s(%s[%d:], %s)", fun, s.bufferName, start, thisField)
		} else {
			return fmt.Sprintf("%s(%s[%s:%s], %s)", fun, s.bufferName, at, end, thisField)
		}
	case "[]byte", "[]uint8":
		//if f.typ != f.aliasType {
		//	fun = f.aliasType
		//}
		if isFirst && isLast {
			if f.Required {
				return fmt.Sprintf("%s(%s[%d:], %s)", fun, s.bufferName, *byteIndex, thisField)
			}
			return fmt.Sprintf("if %s != 0 {\n%s(%s[%d:], %s)\n}", lenVar, fun, s.bufferName, *byteIndex, thisField)
		} else {
			if f.Required {
				return fmt.Sprintf("%s(%s[%s:%s], %s)", fun, s.bufferName, at, end, thisField)
			}
			return fmt.Sprintf("if %s != 0 {%s(%s[%s:%s], %s)\n}", lenVar, fun, s.bufferName, at, end, thisField)
		}
	case "[]int8":
		return fmt.Sprintf("%s(%s[%s:%s], %s)", fun, s.bufferName, at, end, thisField)
	case "[]bool":
		return fmt.Sprintf("%s(%s[%s:%s], %s, %s)", fun, s.bufferName, at, end, thisField, lenVar)
	}

	switch size {
	case 1:
		return fmt.Sprintf("%s[%d]=%s", s.bufferName, start, printFunc(fun, thisField))
	default:
		if start == 0 {
			return printFunc(fun, fmt.Sprintf("%s[:%d]", s.bufferName, *byteIndex), thisField)
		} else {
			if f.isArray() && fun == "copy" {
				return printFunc(fun, fmt.Sprintf("%s[%d:%d]", s.bufferName, start, *byteIndex), thisField+"[:]")
			} else {
				return printFunc(fun, fmt.Sprintf("%s[%d:%d]", s.bufferName, start, *byteIndex), thisField)
			}
		}
	case 0:
		// Variable length size.
		slice := s.bufferName
		idx := at
		if start >= 1 {
			if at == "" {
				slice = fmt.Sprintf("%s[%d:]", s.bufferName, start)
				idx = "0"
			} else if at != "0" {
				slice = fmt.Sprintf("%s[%s:]", s.bufferName, at)
			}
		}

		switch f.typ {
		case "struct":
			return pkgSelName(thisField, printFunc(fun, slice))
		}

		if isLast {
			return printFunc(fun, slice, thisField, "l"+Utoa(index))
		} else {
			return printFunc(fun, slice, thisField, "l"+Utoa(index), idx)
		}
	}

	/*switch f.typ {
	case /*"bool",* / "byte", "uint8":
		return fmt.Sprintf("b[%d]=%s", start, thisField)
	case "int8":
		return fmt.Sprintf("b[%d]=%s", start, o.marshalFunc(f.typ, thisField))
		/*case "int8":
		return fmt.Sprintf("b[%d]=byte(%s)", start, o.marshalFunc(f.typ, thisField))* /
	case "int", "int16", "int32", "int64", "float32", "float64", "uint", "uint16", "uint32", "uint64":
		return o.marshalFunc(f.typ, buffer, thisField)
	case "string":
		slice := "b"
		if at != "0" {
			slice = fmt.Sprintf("b[%v:]", at)
		}
		//return fmt.Sprintf("%s.WriteString(b%s, %s)", pkgName, slice, thisField)
		//return fmt.Sprintf("%s(b%s, %s)", o.marshalFunc(f.typ, thisField), slice, thisField)
		//return fmt.Sprintf("%s", o.marshalFunc(f.typ, slice, thisField, "l1"))
		return o.marshalFunc(f.typ, slice, thisField, "l1")
	}
	return ""*/
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

func printFunc(fun string, params ...string) string {
	if fun == "" {
		return strings.Join(params, ", ")
	}
	b := fmt.Sprintf("%s(%s)", fun, strings.Join(params, ", "))
	return b
}

func (o Option) typeFuncs(fe field, importJ *bool) (fun string, size, totalSize uint) {
	var f interface{}
	switch fe.typ {
	case "byte", "uint8":
		return "", 1, 1
	case "int8":
		return "byte", 1, 1
	case "string":
		return "copy", 0, 0
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				f, size, totalSize = jay.WriteIntArch32, 4, 4
			}
			f, size, totalSize = jay.WriteIntArch64, 8, 8
			break
		}
		f, size, totalSize = jay.WriteIntVariable, 0, 1
	case "int16":
		f, size, totalSize = jay.WriteInt16, 2, 2
	case "int32", "rune":
		f, size, totalSize = jay.WriteInt32, 4, 4
	case "float32":
		f, size, totalSize = jay.WriteFloat32, 4, 4
	case "float64":
		f, size, totalSize = jay.WriteFloat64, 8, 8
	case "int64":
		f, size, totalSize = jay.WriteInt64, 8, 8
	case "time.Duration":
		f, size, totalSize = jay.WriteDuration, 8, 8
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				f, size, totalSize = jay.WriteUintArch32, 4, 4
			}
			f, size, totalSize = jay.WriteUintArch64, 8, 8
			break
		}
		f, size, totalSize = jay.WriteUintVariable, 0, 1
	case "uint16":
		f, size, totalSize = jay.WriteUint16, 2, 2
	case "uint32":
		f, size, totalSize = jay.WriteUint32, 4, 4
	case "uint64":
		f, size, totalSize = jay.WriteUint64, 8, 8
	case "time.Time":
		if fe.tagOptions.TimeNano {
			f, size, totalSize = jay.WriteTimeNano, 8, 8
		} else {
			f, size, totalSize = jay.WriteTime, 8, 8
		}
	case "[]uint8", "[]byte":
		return "copy", 0, 0
	case "[]int8":
		f, size, totalSize = jay.WriteInt8s, 0, 1
	case "[]bool":
		f, size, totalSize = jay.WriteBools, 0, 1

	case "[15]byte", "[15]uint8":
		return "copy", uint(fe.arraySize), uint(fe.arraySize)

	default:
		log.Printf("no function set for type %s yet in typeFuncs()", fe.typ)
		return "", 0, 0
	}

	return nameOf(f, importJ), size, totalSize
}

func nameOf(f any, importJ *bool) string {
	if s, ok := f.(string); ok {
		return s
	}

	if importJ != nil {
		*importJ = true
	}

	return strings.TrimPrefix(
		runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
		pkgPrefix,
	)
}
