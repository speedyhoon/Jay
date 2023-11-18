package jay

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// MakeMarshalJ ...
func (s *Struct) MakeMarshalJ(b *bytes.Buffer) {
	b.WriteString(fmt.Sprintf(
		`func (%s *%s) MarshalJ() (b []byte) {
	b = make([]byte, %[1]s.SizeJ())
	%[1]s.MarshalJTo(b)
	return b
}
`,
		s.ReceiverName(),
		s.name,
	))
}

// MakeMarshalJTo ...
func (s *Struct) MakeMarshalJTo(b *bytes.Buffer) {
	receiver := s.ReceiverName()
	b.WriteString(fmt.Sprintf(
		"func (%s *%s) MarshalJTo(b []byte) {\n",
		receiver,
		s.name,
	))

	var byteIndex uint
	for _, f := range s.fixedLen {
		b.WriteString(generateLine(f, &byteIndex, receiver, ""))
		b.WriteString("\n")
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
				b.WriteString("at:=")
			default:
				b.WriteString("at=")
			}
		}

		b.WriteString(generateLine(f, &byteIndex, receiver, at))
		b.WriteString("\n")
	}

	b.WriteString("}\n")
}
func generateLine(f field, index *uint, receiver, at string) string {
	start := *index
	*index += isLen(f.typ)
	thisField := fmt.Sprintf("%s.%s", receiver, f.name)

	buffer := "b"
	if start >= 1 {
		if at == "" {
			buffer = fmt.Sprintf("b[%d:]", start)
		} else {
			buffer = fmt.Sprintf("b[%s:]", at)
		}
	}

	switch f.typ {
	case "bool", "byte", "int8", "uint8":
		return fmt.Sprintf("b[%d]=%s", start, marshalFunc(f.typ, thisField))
		/*case "int8":
		return fmt.Sprintf("b[%d]=byte(%s)", start, marshalFunc(f.typ, thisField))*/
	case "uint64":
		return marshalFunc(f.typ, thisField, buffer)
	case "string":
		var slice string
		if at != "0" {
			slice = fmt.Sprintf("[%v:]", at)
		}
		return fmt.Sprintf("%s.WriteString(b%s, %s)", pkgName, slice, thisField)

	case "struct":
		return fmt.Sprintf("%s.MarshalJTo(%s)", thisField, buffer)

	case "int", "uint":
		return marshalFunc(f.typ, thisField, buffer)
	default:
		log.Printf("no generateLine for type `%s` yet", f.typ)
	}
	return ""
}

/*func lookupMarshaller(f *field) bool {
	switch f.typ {
	case "uint64":
		marshalU64(f)
		return true
	}
	return false
}*/

/*func marshalU64(f *field) (fun string, sizeOf uint, ok bool) {
	if f.tagOptions.maxBytes != 0 {
		return fmt.Sprintf("WriteUint%dBytes", f.tagOptions.maxBytes*8), f.tagOptions.maxBytes, true
	}
	return "", 0, false
}*/

func Utoa(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

func marshalFunc(typ, field string, params ...string) string {
	switch typ {
	case "byte", "uint8":
		return field
	case "int8":
		return fmt.Sprintf("int8(%s)", field)
	default:
		if len(params) >= 1 {
			field = strings.Join(append(params, field), ", ")
		}

		return fmt.Sprintf("%s(%s)", getFuncName(typeFuncs(typ)), field)
	}
}

func getFuncName(f interface{}) string {
	if f == nil {
		return ""
	}
	t := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return strings.TrimPrefix(
		t,
		importPrefix,
	)
}

func typeFuncs(typ string) interface{} {
	switch typ {
	// TODO "byte", "int8", "uint8", "float32","float64",
	case "bool":
		return Bool1
	case "string":
		return WriteString
	case "uint":
		return WriteUint
	case "uint16":
		return WriteUint16
	case "uint32":
		return WriteUint32
	case "uint64":
		return WriteUint64
	case "int":
		return WriteInt
	case "int16":
		return WriteInt16
	case "int32", "rune":
		return WriteInt32
	case "int64":
		return WriteInt64
	case "byte", "int8", "uint8":
		// None needed.
	default:
		log.Printf("not function set for type %s yet", typ)
	}
	return nil
}

/*
func (e *Event) MarshalZTo(buf []byte) {
	//Bool1(buf[:1], e.Auto)
	//Bool1(buf[1:2], e.MOT)
	//WriteBoo2(buf[0], e.Auto)
	//WriteBoo2(buf[1], e.MOT)
	//buf[0] = WriteBoo3(e.Auto)
	//buf[1] = WriteBoo3(e.MOT)

	buf[0] = WriteBoo5(e.Auto, e.MOT, e.ABS, e.TCS)
	buf[1] = WriteEnum44(e.Enum1, e.Enum2)

	at := 10
	WriteUint64(buf[1:at], e.ID) // 0-7
	at = WriteString(buf, at, &e.Name)
	at = WriteString(buf, at, &e.CC)
	WriteString(buf, at, &e.Timing)
}
*/
