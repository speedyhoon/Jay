package generate

import (
	"log"
	"strconv"
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
	if l := len(s.bool); l >= 1 {
		qty = boolsSliceIndex(uint(l)) + 1
	}

	qty += uint(len(s.single))

	for _, x := range s.fixedLen {
		qty += o.typeFuncSize(x.typ)
	}
	for _, v := range s.variableLen {
		qty += o.typeFuncSize(v.typ)
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
func (o Option) typeFuncSize(typ string) (size uint) {
	switch typ {
	case "string", "bool", "byte", "uint8", "int8", "[]byte":
		return 1
	case "int16", "uint16":
		return 2
	case "int32", "rune", "float32", "uint32":
		return 4
	case "float64", "int64", "uint64", "time.Time":
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

	default:
		if typeArray.MatchString(typ) {
			underlyingType := typeArrayBrackets.ReplaceAllString(typ, "")

			itemSize := o.typeFuncSize(underlyingType)
			sizeStr := typ[1 : len(typ)-1-len(underlyingType)]
			siz, err := strconv.ParseUint(sizeStr, 10, 0)
			if err != nil {
				log.Printf("array %s size `%s` is invalid", typ, sizeStr)
				return 0
			}
			return uint(siz) * itemSize
		}

		log.Printf("no function set yet for type %s in typeFuncSize()", typ)
		return 0
	}
}
