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

// MakeMarshalJX ...
func (s *Struct) MakeMarshalJX(b *bytes.Buffer, o Option) {
	receiver := s.ReceiverName()
	b.WriteString(fmt.Sprintf(
		"func (%s *%s) MarshalJX() {\n%s\nb := make([]byte, %s)\n",
		receiver,
		s.name,
		lengths2(s.varLenFieldNames(), receiver),
		joinSizes(s.calcSize(o), s.variableLen, s.ReceiverName()),
	))

	var byteIndex uint
	if len(s.bool) != 0 {
		s.generateBools(s.bool, b, o, &byteIndex, receiver)
	}

	for _, f := range s.fixedLen {
		b.WriteString(o.generateLine(f, &byteIndex, receiver, "", 0))
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

		b.WriteString(o.generateLine(f, &byteIndex, receiver, at, uint(i)))
		b.WriteString("\n")
	}

	b.WriteString("}\n")
}

func (s *Struct) generateBools(bools []field, b *bytes.Buffer, o Option, byteIndex *uint, receiver string) {
	var i, l uint = 0, uint(len(s.bool) / 8)
	for ; i <= l; i++ {
		WriteBools(bools[BoolsSliceIndex(i):], b, o, byteIndex, receiver)
	}
}

func BoolsSliceIndex(input uint) uint {
	if input == 0 {
		return 0
	}
	return ((input-1)/8+1)*8 - 8
}

func WriteBools(bools []field, b *bytes.Buffer, o Option, byteIndex *uint, receiver string) {
	if len(bools) > 8 {
		bools = bools[:8]
	}

	b.WriteString(fmt.Sprintf("b[%d] = %s.Bool%d(%s)\n", *byteIndex, o.pkgName, len(bools), fieldNames(bools, receiver)))

	*byteIndex++
}
func fieldNames(fields []field, receiver string) string {
	var s []string
	for i := range fields {
		s = append(s, fmt.Sprintf("%s.%s", receiver, fields[i].name))
	}
	return strings.Join(s, ", ")
}

// MakeMarshalJTo ...
func (s *Struct) MakeMarshalJTo(o Option, b *bytes.Buffer) {
	receiver := s.ReceiverName()
	b.WriteString(fmt.Sprintf(
		"func (%s *%s) MarshalJTo(b []byte) {\n%s",
		receiver,
		s.name,
		lengths2(s.varLenFieldNames(), receiver),
	))

	var byteIndex uint
	if len(s.bool) != 0 {
		s.generateBools(s.bool, b, o, &byteIndex, receiver)
	}

	for _, f := range s.fixedLen {
		b.WriteString(o.generateLine(f, &byteIndex, receiver, "", 0))
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

		b.WriteString(o.generateLine(f, &byteIndex, receiver, at, uint(i)))
		b.WriteString("\n")
	}

	b.WriteString("}\n")
}

func (o Option) generateLine(f field, byteIndex *uint, receiver, at string, index uint) string {
	fun, size := o.typeFuncs(f.typ)
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

		//at := 12
		//	at = jay.WriteStringN(b[at:at+l1+1], c.Name, l1, at)

		return printFunc(fun, slice, thisField, "l"+Utoa(index), idx)
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

	case "struct":
		return fmt.Sprintf("%s.MarshalJTo(%s)", thisField, buffer)
	}
	return ""*/
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

func printFunc(fun string, params ...string) string {
	if fun == "" {
		return strings.Join(params, ", ")
	}
	b := fmt.Sprintf("%s(%s)" /*getFuncName(*/, fun /*)*/, strings.Join(params, ", "))
	return b
}

/*func (o Option) marshalFunc(typ string, params ...string) string {
	//switch typ {
	//case "byte", "uint8":
	//	return strings.Join(params, ", ")
	//default:
	return fmt.Sprintf("%s(%s)", getFuncName(o.typeFuncs(typ)), strings.Join(params, ", "))
	//}
}*/

/*func getFuncName(f interface{}) string {
	switch s := f.(type) {
	case string:
		return s
	case nil:
		return ""
	default:
		return strings.TrimPrefix(
			runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
			importPrefix, // TODO replace with: filepath.Dir(get package path during run time)
		)
	}
}*/

func (o Option) typeFuncs(typ string) (_ string, size uint) {
	var f interface{}
	switch typ {
	case "byte", "uint8":
		return "", 1
	case "int8":
		return "byte", 1
	case "bool":
		f, size = Bool1, 1
	case "string":
		f, size = WriteStringN, 0
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				f, size = WriteIntArch32, 4
			}
			f, size = WriteIntArch64, 8
			break
		}
		f, size = WriteIntVariable, 0
	case "int16":
		f, size = WriteInt16, 2
	case "int32", "rune":
		f, size = WriteInt32, 4
	case "float32":
		f, size = WriteFloat32, 4
	case "float64":
		f, size = WriteFloat64, 8
	case "int64":
		f, size = WriteInt64, 8
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				f, size = WriteUintArch32, 4
			}
			f, size = WriteUintArch64, 8
			break
		}
		f, size = WriteUintVariable, 0
	case "uint16":
		f, size = WriteUint16, 2
	case "uint32":
		f, size = WriteUint32, 4
	case "uint64":
		f, size = WriteUint64, 8
	case "struct":
		return "MarshalJTo", 0

	default:
		log.Printf("not function set for type %s yet", typ)
		return "", 0
	}

	return strings.TrimPrefix(
		runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
		o.importPrefix,
	), size
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
