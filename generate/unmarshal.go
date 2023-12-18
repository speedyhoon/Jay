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

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer, o Option) {
	receiver := s.receiverName()
	var byteIndex uint
	buf := bytes.NewBuffer(nil)

	s.makeReadBools(buf, &byteIndex, receiver)

	for i, f := range s.fixedLen {
		buf.WriteString(o.unmarshalLine(f, &byteIndex, receiver, "", i == len(s.fixedLen)-1 && len(s.variableLen) == 0))
		buf.WriteString("\n")
	}

	var at string
	if len(s.variableLen) >= 1 {
		buf.WriteString("var ok bool\n")
	}

	if len(s.variableLen) == 1 {
		at = strconv.Itoa(len(s.variableLen))
	} else if len(s.variableLen) >= 2 {
		bufWriteF(buf, "at:=%d\n", byteIndex)
		at = "at"
	}

	vLen := len(s.variableLen) - 1
	for i, f := range s.variableLen {
		buf.WriteString(o.unmarshalLine(f, &byteIndex, receiver, at, i == vLen))
		bufWriteF(buf, "\nif !ok {\nreturn %s.ErrUnexpectedEOB\n}\n", pkgName)
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	bufWriteF(b,
		"func (%s *%s) UnmarshalJ(b []byte) error {\n%s\nreturn nil\n}\n",
		receiver,
		s.name,
		code,
	)
}

func (o Option) unmarshalLine(f field, byteIndex *uint, receiver, at string, isLast bool) string {
	fun, size := o.unmarshalFuncs(f, isLast)
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

// unmarshalFuncs returns the function name to handle unmarshalling.
// `size` is the quantity of bytes required to represent the type.
func (o Option) unmarshalFuncs(f field, isLast bool) (funcName string, size uint) {
	var c interface{}
	switch f.typ {
	case "byte", "uint8":
		return "", 1
	case "int8":
		return "int8", 1
	case "string":
		if isLast {
			c, size = jay.ReadString, 0
		} else {
			c, size = jay.ReadStringAt, 0
		}
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				c, size = jay.ReadIntArch32, 4
			}
			c, size = jay.ReadIntArch64, 8
			break
		}
		//c, size = jay.ReadIntVariable, 0
		c, size = jay.ReadInt, 0
	case "int16":
		c, size = jay.ReadInt16, 2
	case "int32", "rune":
		c, size = jay.ReadInt32, 4
	case "float32":
		c, size = jay.ReadFloat32, 4
	case "float64":
		c, size = jay.ReadFloat64, 8
	case "int64":
		c, size = jay.ReadInt64, 8
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				c, size = jay.ReadUintArch32, 4
			}
			c, size = jay.ReadUintArch64, 8
			break
		}
		c, size = jay.ReadUintVariable, 0
	case "uint16":
		c, size = jay.ReadUint16, 2
	case "uint32":
		c, size = jay.ReadUint32, 4
	case "uint64":
		c, size = jay.ReadUint64, 8
	case "time.Time":
		if f.tagOptions.TimeNano {
			c, size = jay.ReadTimeNano, 8
		} else {
			c, size = jay.ReadTime, 8
		}

	default:
		log.Printf("no function set for type %s yet", f.typ)
		return "", 0
	}

	return strings.TrimPrefix(
		runtime.FuncForPC(reflect.ValueOf(c).Pointer()).Name(),
		pkgPrefix,
	), size
}
