package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// TODO add support for pointers with all types.
// TODO add support for enums with restricted sizes like: buf[1] = WriteEnum44(e.Enum1, e.Enum2)

// makeMarshal ...
func (s *structTyp) makeMarshal(b *bytes.Buffer, o Option) {
	varLengths := lengths2(s.varLenFieldNames(o), s.receiver)
	makeSize := joinSizes(s.calcSize(o), s.variableLen, o)

	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)
	isReturnInlined := s.makeWriteBools(buf, &byteIndex)
	isReturnInlined = s.writeSingles(buf, &byteIndex, s.receiver, o) || isReturnInlined

	for i, f := range s.fixedLen {
		buf.WriteString(o.generateLine(f, &byteIndex, s.receiver, "", 0, i == len(s.fixedLen)-1 && len(s.variableLen) == 0, s.bufferName))
		buf.WriteString("\n")
	}

	at := Utoa(byteIndex)
	vLen := len(s.variableLen) - 1
	for i, f := range s.variableLen {
		if i == 1 {
			at = "at"
		}
		if i != vLen {
			switch i {
			case 0:
				buf.WriteString("at:=")
			default:
				buf.WriteString("at=")
			}
		}

		buf.WriteString(o.generateLine(f, &byteIndex, s.receiver, at, uint(i), i == vLen, s.bufferName))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	if isReturnInlined {
		bufWriteF(b,
			"func (%s *%s) MarshalJ() []byte {\nreturn []byte{%s}\n}\n",
			s.receiver,
			s.name,
			code,
		)
		return
	}

	b.WriteString(fmt.Sprintf(
		"func (%[1]s *%[2]s) MarshalJ() (%[3]s []byte) {\n%[4]s\n%[3]s = make([]byte, %[5]s)\n%[6]s%[7]sreturn\n}\n",
		s.receiver,
		s.name,
		s.bufferName,
		varLengths,
		makeSize,
		s.generateSizeLine(),
		code,
	))
}

func fieldNames(fields []field, receiver string) string {
	var s []string
	for i := range fields {
		s = append(s, fmt.Sprintf("%s.%s", receiver, fields[i].name))
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
		values[i] = fmt.Sprintf("byte(l%d)", i)
	}
	return fmt.Sprintln(strings.Join(assignments, ", "), " = ", strings.Join(values, ", "))
}

func (o Option) generateLine(f field, byteIndex *uint, receiver, at string, index uint, isLast bool, bufferName string) string {
	fun, size, totalSize := o.typeFuncs(f, isLast)
	if fun == "" && size == 0 {
		// Unknown type, not supported yet.
		log.Printf("no generateLine for type `%s` yet", f.typ)
		return ""
	}

	start := *byteIndex
	*byteIndex += totalSize
	thisField := fmt.Sprintf("%s.%s", receiver, f.name)

	switch size {
	case 1:
		return fmt.Sprintf("%s[%d]=%s", bufferName, start, printFunc(fun, thisField))
	default:
		if start == 0 {
			return printFunc(fun, fmt.Sprintf("%s[:%d]", bufferName, *byteIndex), thisField)
		} else {
			if f.isArray() && fun == "copy" {
				return printFunc(fun, fmt.Sprintf("%s[%d:%d]", bufferName, start, *byteIndex), thisField+"[:]")
			} else {
				return printFunc(fun, fmt.Sprintf("%s[%d:%d]", bufferName, start, *byteIndex), thisField)
			}
		}
	case 0:
		// Variable length size.
		slice := bufferName
		idx := at
		if start >= 1 {
			if at == "" {
				slice = fmt.Sprintf("%s[%d:]", bufferName, start)
				idx = "0"
			} else if at != "0" {
				slice = fmt.Sprintf("%s[%s:]", bufferName, at)
			}
		}

		switch f.typ {
		case "struct":
			return fmt.Sprintf("%s.%s", thisField, printFunc(fun, slice))
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

func (o Option) typeFuncs(fe field, isLast bool) (_ string, size, totalSize uint) {
	var f interface{}
	switch fe.typ {
	case "byte", "uint8":
		return "", 1, 1
	case "int8":
		return "byte", 1, 1
	case "string":
		if isLast {
			f, size, totalSize = jay.WriteString, 0, 1
		} else {
			f, size, totalSize = jay.WriteStringAt, 0, 1
		}
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
	case "[]byte":
		if isLast {
			f, size, totalSize = jay.WriteBytes, 0, 1
		} else {
			f, size, totalSize = jay.WriteBytesAt, 0, 1
		}

	case "[15]byte", "[15]uint8":
		return "copy", uint(fe.arraySize), uint(fe.arraySize)

	default:
		log.Printf("no function set for type %s yet in typeFuncs()", fe.typ)
		return "", 0, 0
	}

	return nameOf(f), size, totalSize
}

func nameOf(f any) string {
	return strings.TrimPrefix(
		runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
		pkgPrefix,
	)
}
