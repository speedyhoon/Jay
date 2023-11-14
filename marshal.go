package jay

import (
	"bytes"
	"fmt"
	"log"
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
	//hasVariable := len(s.variableLen) != 0
	//var isLast bool
	for _, f := range s.fixedLen {
		//if !hasVariable {
		//	isLast = i >= len(s.fixedLen)-1
		//}

		//if lookupMarshaller(&f) {
		//	continue
		//}

		b.WriteString(generateLine(f, &byteIndex, receiver, "" /*, isLast*/))
		b.WriteString("\n")
	}

	at := Utoa(byteIndex)
	vLen := len(s.variableLen) - 1
	for i, f := range s.variableLen {
		switch i {
		case 0:
			b.WriteString("at:=")
		case 1:
			at = "at"
			b.WriteString("at=")
		case vLen:
		default:
			b.WriteString("at=")
		}

		//isLast = i >= vLen
		b.WriteString(generateLine(f, &byteIndex, receiver, at /*, isLast*/))
		b.WriteString("\n")
	}

	b.WriteString("}\n")
}
func generateLine(f field, index *uint, receiver, at string /*, isLast bool*/) string {
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
	case "uint64":
		return marshalFunc(f.typ, thisField, buffer)
	case "string":
		//if !isLast {
		// TODO thisField should == b[1:]
		return fmt.Sprintf("%sWriteString(b, %s, &%s)", packageName, at, thisField)
	//} else {
	//	return fmt.Sprintf("%sWriteString(b, %s, &%s.%s)", packageName, at, receiver, f.name)
	//}

	case "struct":
		return fmt.Sprintf("%s.MarshalJTo(%s)", thisField, buffer)

	case "int":
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
	if f.Options.maxBytes != 0 {
		return fmt.Sprintf("WriteUint%dBytes", f.Options.maxBytes*8), f.Options.maxBytes, true
	}
	return "", 0, false
}*/

func (f *field) LoadTagOptions() {
	f.tag = strings.TrimSpace(f.tag)
	if f.tag == "" {
		return
	}
	for _, c := range strings.Split(f.tag, ",") {
		d := strings.Split(c, ":")
		switch g := d[0]; g {
		case max:
			loadUint(d[1], &f.Options.Max)
			f.Options.maxBytes = byteSize(f.Options.Max)
		case min:
			loadUint(g, &f.Options.Min)
		}
	}
}

func byteSize(v uint) uint {
	if v == 0 {
		return 0
	}
	if v <= MaxUint8 {
		return 1
	}
	if v <= MaxUint16 {
		return 2
	}
	if v <= MaxUint24 {
		return 3
	}
	if v <= MaxUint32 {
		return 4
	}
	if v <= MaxUint40 {
		return 5
	}
	if v <= MaxUint48 {
		return 6
	}
	if v <= MaxUint56 {
		return 7
	}
	return 8
}

const (
	max = "max"
	min = "min"
)

func loadUint(g string, f *uint) {
	uu, err := strconv.ParseUint(g, 10, 64)
	if err != nil {
		log.Println(err)
	}
	*f = uint(uu)
}

type Options struct {
	Max      uint
	maxBytes uint
	Min      uint
}

func Utoa(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

const packageName = "jay."

func marshalFunc(typ, field string, params ...string) string {
	switch typ {
	case "byte", "int8", "uint8":
		return field
	default:
		if len(params) >= 1 {
			field = strings.Join(append(params, field), ", ")
		}

		return fmt.Sprintf("%s%s(%s)", packageName, typeFuncs(typ), field)
	}
}
func typeFuncs(typ string) string {
	switch typ {
	// TODO "float32","float64",
	//case "byte", "int8", "uint8", "string", "bool":
	//	return 1
	//case "int16", "uint16":
	//	return 2
	//case "int32", "rune", "uint32":
	//	return 4
	//case "int", "int64", "uint", "uint64":
	//	return 8
	case "bool":
		return "Bool1"
	case "string":
		return "WriteString"
	case "uint64":
		return "WriteUint64"
	case "int":
		return "WriteInt"
	default:
		log.Printf("not func set for type %s yet", typ)
	}
	return ""
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
