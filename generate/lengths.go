package generate

import (
	"fmt"
	"strconv"
	"strings"
)

/*// LenDecl ...
func (s *structTyp) LenDecl(b *bytes.Buffer) {
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
		lengths2(variableFields, s.receiverName()),
		structs2(variableStructs, s.receiverName()),
	))
}*/

func joinSizes(qty uint, variableLen []field, o Option) string {
	var s []string
	if qty != 0 {
		s = []string{Utoa(qty)}
	}

	var variableFields, variableStructs []string
	for _, v := range variableLen {
		qty += isLen(v.typ)
		if o.isLenVariable(v.typ) {
			variableFields = append(variableFields, v.name)
		} else if v.typ == "struct" {
			variableStructs = append(variableStructs, v.name)
		}
	}

	if len(variableFields) != 0 {
		s = append(s, addDecls(len(variableFields)))
	}
	//if len(variableStructs) != 0 {
	//	s = append(s, lengths2(variableStructs, receiverName))
	//}

	return strings.Join(s, "+")
}

func (s *structTyp) varLenFieldNames(o Option) (names []string) {
	for _, v := range s.variableLen {
		if o.isLenVariable(v.typ) {
			names = append(names, v.name)
			//} else if v.typ == "struct" {
			//variableStructs = append(variableStructs, v.name)
		}
	}
	return
}

func lengths2(names []string, receiver string) string {
	if len(names) == 0 {
		return ""
	}

	receiver = fmt.Sprintf("len(%s.", receiver)
	declarations := strings.Join(decls(len(names)), ", ")
	return fmt.Sprintf("%s := %s%s)",
		declarations,
		receiver,
		strings.Join(names, "),"+receiver),
	)
}

func decls(u int) (s []string) {
	for i := 0; i < u; i++ {
		s = append(s, "l"+strconv.Itoa(i))
	}
	return
}

func addDecls(u int) string {
	return strings.Join(decls(u), "+")
}

/*func structs2(names []string, receiver string) string {
	if len(names) == 0 {
		return ""
	}

	return fmt.Sprintf(
		"+ %s.%s.SizeJ()",
		receiver,
		strings.Join(names, ".SizeJ()+"+receiver+"."),
	)
}*/
