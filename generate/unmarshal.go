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

// MakeUnmarshal ...
func (s *structTyp) MakeUnmarshal(b *bytes.Buffer, o Option) {
	receiver := s.ReceiverName()
	var byteIndex uint
	buf := bytes.NewBuffer(nil)

	s.generateReadBools(buf, &byteIndex, receiver)

	for i, f := range s.fixedLen {
		buf.WriteString(o.unmarshalLine(f, &byteIndex, receiver, "", 0, i == len(s.fixedLen)-1 && len(s.variableLen) == 0))
		buf.WriteString("\n")
	}

	var at string
	if len(s.variableLen) >= 1 {
		buf.WriteString("var ok bool\n")
	}

	if len(s.variableLen) == 1 {
		at = strconv.Itoa(len(s.variableLen))
	} else if len(s.variableLen) >= 2 {
		buf.WriteString(fmt.Sprintf("at:=%d\n", byteIndex))
		at = "at"
	}

	vLen := len(s.variableLen) - 1
	for i, f := range s.variableLen {
		buf.WriteString(o.unmarshalLine(f, &byteIndex, receiver, at, uint(i), i == vLen))
		buf.WriteString(fmt.Sprintf("\nif !ok {\nreturn %s.ErrUnexpectedEOB\n}\n", pkgName))
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	b.WriteString(fmt.Sprintf(
		"func (%s *%s) UnmarshalJ(b []byte) (err error) {\n%s\nreturn nil\n}\n",
		receiver,
		s.name,
		code,
	))
}

func (o Option) unmarshalLine(f field, byteIndex *uint, receiver, at string, index uint, isLast bool) string {
	fun, size := o.unmarshalFuncs(f.typ, isLast)
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
		return fmt.Sprintf("%s=%s", thisField, printFunc(fun, fmt.Sprintf("b[%d]", start)))
	default:
		return fmt.Sprintf("%s = %s", thisField, printFunc(fun, fmt.Sprintf("b[%d:%d]", start, *byteIndex)))
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
			return fmt.Sprintf("%s, at, ok = %s", thisField, printFunc(fun, slice, idx))
		} else {
			return fmt.Sprintf("%s, at, ok = %s", thisField, printFunc(fun, slice))
		}
	}
}

func (o Option) unmarshalFuncs(typ string, isLast bool) (_ string, size uint) {
	var f interface{}
	switch typ {
	case "byte", "uint8":
		return "", 1
	case "int8":
		return "int8", 1
	case "string":
		if isLast {
			f, size = jay.ReadString, 0
		} else {
			f, size = jay.ReadStringAt, 0
		}
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				f, size = jay.ReadIntArch32, 4
			}
			f, size = jay.ReadIntArch64, 8
			break
		}
		//f, size = jay.ReadIntVariable, 0
		f, size = jay.ReadInt, 0
	case "int16":
		f, size = jay.ReadInt16, 2
	case "int32", "rune":
		f, size = jay.ReadInt32, 4
	case "float32":
		f, size = jay.ReadFloat32, 4
	case "float64":
		f, size = jay.ReadFloat64, 8
	case "int64":
		f, size = jay.ReadInt64, 8
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				f, size = jay.ReadUintArch32, 4
			}
			f, size = jay.ReadUintArch64, 8
			break
		}
		f, size = jay.ReadUintVariable, 0
	case "uint16":
		f, size = jay.ReadUint16, 2
	case "uint32":
		f, size = jay.ReadUint32, 4
	case "uint64":
		f, size = jay.ReadUint64, 8

	default:
		log.Printf("not function set for type %s yet", typ)
		return "", 0
	}

	return strings.TrimPrefix(
		runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
		pkgPrefix,
	), size
}
