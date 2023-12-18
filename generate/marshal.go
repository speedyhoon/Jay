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
	receiver := s.receiverName()

	varLengths := lengths2(s.varLenFieldNames(), receiver)

	var byteIndex uint
	buf := bytes.NewBuffer(nil)
	s.makeWriteBools(buf, &byteIndex, receiver)

	for i, f := range s.fixedLen {
		buf.WriteString(o.generateLine(f, &byteIndex, receiver, "", 0, i == len(s.fixedLen)-1 && len(s.variableLen) == 0))
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

		buf.WriteString(o.generateLine(f, &byteIndex, receiver, at, uint(i), i == vLen))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	bufWriteF(b,
		"func (%s *%s) MarshalJ() (b []byte) {\n%s\nb = make([]byte, %s)\n%sreturn\n}\n",
		receiver,
		s.name,
		varLengths,
		joinSizes(s.calcSize(o), s.variableLen),
		code,
	)
}

func fieldNames(fields []field, receiver string) string {
	var s []string
	for i := range fields {
		s = append(s, fmt.Sprintf("%s.%s", receiver, fields[i].name))
	}
	return strings.Join(s, ", ")
}

func (o Option) generateLine(f field, byteIndex *uint, receiver, at string, index uint, isLast bool) string {
	fun, size := o.typeFuncs(f, isLast)
	if fun == "" && size == 0 {
		// Unknown type, not supported yet.
		log.Printf("no generateLine for type `%s` yet", f.typ)
		return ""
	}

	start := *byteIndex
	*byteIndex += size
	thisField := fmt.Sprintf("%s.%s", receiver, f.name)

	switch size {
	case 1:
		return fmt.Sprintf("b[%d]=%s", start, printFunc(fun, thisField))
	default:
		return printFunc(fun, fmt.Sprintf("b[%d:%d]", start, *byteIndex), thisField)
	case 0:
		// Variable length size.
		slice := "b"
		idx := at
		if start >= 1 {
			if at == "" {
				slice = fmt.Sprintf("b[%d:]", start)
				idx = "0"
			} else if at != "0" {
				slice = fmt.Sprintf("b[%s:]", at)
			}
		}

		switch f.typ {
		case "struct":
			return fmt.Sprintf("%s.%s", thisField, printFunc(fun, slice))
		}

		if !isLast {
			return printFunc(fun, slice, thisField, "l"+Utoa(index), idx)
		} else {
			return printFunc(fun, slice, thisField, "l"+Utoa(index))
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

func (o Option) typeFuncs(fe field, isLast bool) (_ string, size uint) {
	var f interface{}
	switch fe.typ {
	case "byte", "uint8":
		return "", 1
	case "int8":
		return "byte", 1
	case "string":
		if isLast {
			f, size = jay.WriteString, 0
		} else {
			f, size = jay.WriteStringAt, 0
		}
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				f, size = jay.WriteIntArch32, 4
			}
			f, size = jay.WriteIntArch64, 8
			break
		}
		f, size = jay.WriteIntVariable, 0
	case "int16":
		f, size = jay.WriteInt16, 2
	case "int32", "rune":
		f, size = jay.WriteInt32, 4
	case "float32":
		f, size = jay.WriteFloat32, 4
	case "float64":
		f, size = jay.WriteFloat64, 8
	case "int64":
		f, size = jay.WriteInt64, 8
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				f, size = jay.WriteUintArch32, 4
			}
			f, size = jay.WriteUintArch64, 8
			break
		}
		f, size = jay.WriteUintVariable, 0
	case "uint16":
		f, size = jay.WriteUint16, 2
	case "uint32":
		f, size = jay.WriteUint32, 4
	case "uint64":
		f, size = jay.WriteUint64, 8
	case "time.Time":
		if fe.tagOptions.TimeNano {
			f, size = jay.WriteTimeNano3, 8
		} else {
			f, size = jay.WriteTime, 8
		}

	default:
		log.Printf("no function set for type %s yet", fe.typ)
		return "", 0
	}

	return strings.TrimPrefix(
		runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
		pkgPrefix,
	), size
}
