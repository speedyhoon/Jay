package jay

import (
	"bytes"
	"fmt"
	"strings"
)

// MakeSize ...
func (s *Struct) MakeSize(b *bytes.Buffer) {
	//b.WriteString(fmt.Sprintf(
	//	"func (%s *%s) SizeJ() int {\n// fields: %d\n}\n",
	//	s.ReceiverName(),
	//	s.name,
	//	len(s.fixedLen),
	//))

	var qty uint
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
	b.WriteString(fmt.Sprintf(
		"func (%s *%s) SizeJ() int {\nreturn %d%s%s\n}\n",
		s.ReceiverName(),
		s.name,
		qty,
		lengths(variableFields, s.ReceiverName()),
		structs(variableStructs, s.ReceiverName()),
	))
}

func lengths(names []string, receiver string) string {
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
}
