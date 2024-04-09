package generate

import (
	"github.com/speedyhoon/jay"
)

/*// MakeSize ...
func (s *structTyp) MakeSize(b *bytes.Buffer) {
	var qty uint
	// TODO add bool sizes
	var variableStructs []string
	for _, x := range s.fixedLen {
		qty += isLen(x.typ)
		if x.typ == "struct" {
			variableStructs = append(variableStructs, x.name)
		}
	}
	var variableFields []string
	for _, v := range s.variableLen {
		qty += isLen(v.typ)
		if v.typ == "string" { //} isLenVariable(v.typ) {
			variableFields = append(variableFields, v.name)
		} else if v.typ == "struct" {
			variableStructs = append(variableStructs, v.name)
		}
	}
	bufWriteF(b,
		"func (%s *%s) SizeJ() int {\nreturn %d%s%s\n}\n",
		s.receiverName(),
		s.name,
		qty,
		lengths(variableFields, s.receiverName()),
		structs(variableStructs, s.receiverName()),
	))
}*/

func (s *structTyp) calcSize(o Option) (qty uint) {
	qty = uint(jay.SizeBools(len(s.bool)))

	qty += uint(len(s.single))

	for _, x := range s.fixedLen {
		qty += x.typeFuncSize(o)
	}
	for _, v := range s.variableLen {
		qty += v.typeFuncSize(o)
	}
	return qty
}

/*func lengths(names []string, receiver string) string {
	if len(names) == 0 {
		return ""
	}
	receiver = fmt.Sprintf("+len(%s.", receiver)
	return receiver + strings.Join(names, ")"+receiver) + ")"
}

func structs(names []string, receiver string) string {
	if len(names) == 0 {
		return ""
	}

	//receiver = fmt.Sprintf(".SizeJ()+%s.", receiver)
	return "+" + receiver + "." + strings.Join(names, ".SizeJ()+"+receiver+".") + ".SizeJ()"
	//+c.Wheels.SizeJ()+c.Gearbox.SizeJ()+c.Roof.SizeJ()
}*/

// typeFuncSize returns the minimum quantity of bytes required to represent an empty or undefined value.
func (f field) typeFuncSize(o Option) (size uint) {
	switch {
	case f.arraySize <= typeSlice:
		return 1
	case f.arraySize >= typeArray:
		itemSize := field{typ: f.arrayType}.typeFuncSize(o)
		return uint(f.arraySize) * itemSize
	case f.arraySize == typeNotArrayOrSlice:
		switch f.typ {
		case "string", "bool", "byte", "uint8", "int8":
			return 1
		case "int16", "uint16":
			return 2
		case "int32", "rune", "float32", "uint32":
			return 4
		case "float64", "int64", "uint64", "time.Time", "time.Duration":
			return 8
		case "int":
			if o.FixedIntSize {
				if o.Is32bit {
					return 4
				}
				return 8
			}
			return 1
		case "uint":
			if o.FixedUintSize {
				if o.Is32bit {
					return 4
				}
				return 8
			}
			return 1
		}
	}
	lg.Printf("type %s unhandled in typeFuncSize()", f.typ)
	return
}
