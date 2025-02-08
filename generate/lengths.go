package generate

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"strconv"
	"strings"
)

func joinSizes(qty uint, variableLen, strSlices []field, importJ *bool) string {
	var s []string
	if qty != 0 {
		s = []string{utl.UtoA(qty)}
	}

	for i, v := range variableLen {
		qty += isLen(v.typ)
		if !v.isFixedLen {
			if v.typ == tBoolS {
				s = append(s, printFunc(nameOf(jay.SizeBools, importJ), lenVariable(i)))
			} else {
				s = append(s, multiples(v, lenVariable(i)))
			}
		}
	}

	vl, sl := len(variableLen), len(strSlices)
	const funcName = "StringsSize16"

	for i, slice := range strSlices {
		if i+1 == sl {
			s = append(s, fmt.Sprintf("%s.%s(%s.%s)", pkgName, funcName, slice.structTyp.receiver, slice.name))
		} else {
			s = append(s, lenVariable(i+vl))
		}
	}

	return strings.Join(s, "+")
}

func multiples(f field, lenVar string) string {
	switch {
	case f.isSlice():
		itemSize := field{typ: f.arrayType, structTyp: f.structTyp}.typeFuncSize()
		if itemSize >= 2 {
			return fmt.Sprintf("%s*%d", lenVar, itemSize)
		}
		return lenVar
	case f.isArray():
		itemSize := field{typ: f.arrayType, structTyp: f.structTyp}.typeFuncSize()
		return utl.UtoA(uint(f.arraySize) * itemSize)
	default: // typeNotArrayOrSlice
		return lenVar
	}
}

func (s *structTyp) varLenFieldNames() (names []string) {
	for _, v := range s.variableLen {
		names = append(names, v.name)
	}
	return
}

func lengths2(names []string, receiver, lenName string) string {
	if len(names) == 0 {
		return ""
	}

	receiver = fmt.Sprintf("len(%s.", receiver)
	declarations := strings.Join(decls(len(names), lenName), ", ")
	return fmt.Sprintf("%s := %s%s)",
		declarations,
		receiver,
		strings.Join(names, "),"+receiver),
	)
}

func lengths3(names, strSlices []string, receiver, lenName string) string {
	qty := len(names) + len(strSlices)
	if qty == 0 {
		return ""
	}

	receiver = fmt.Sprintf("len(%s.", receiver)
	declarations := strings.Join(decls(qty, lenName), ", ")
	return fmt.Sprintf("%s := %s%s)",
		declarations,
		receiver,
		strings.Join(names, "),"+receiver),
	)
}

func (s *structTyp) defineLengths() (out string) {
	vl, sl := len(s.variableLen), len(s.stringSlice)-1

	qty := vl + sl
	if qty <= 0 {
		return ""
	}

	if vl >= 1 {
		receiver := fmt.Sprintf("len(%s.", s.receiver)
		out = receiver + strings.Join(s.varLenFieldNames(), "),"+receiver) + ")"
		if sl >= 1 {
			out += ","
		}
	}

	if sl >= 1 {
		names := make([]string, sl)

		const funcName = "StringsSize16"
		for i, strSlice := range s.stringSlice[:sl] {
			names[i] = fmt.Sprintf("%s.%s(%s.%s)", pkgName, funcName, s.receiver, strSlice.name)
		}

		out += strings.Join(names, ", ")
	}

	declarations := strings.Join(decls(qty, s.lengthName), ", ")
	return fmt.Sprintf("%s := %s", declarations, out)
}

func assignStringSliceSizes(fl fieldList) (_ string, declarations []string) {
	names := make([]string, len(fl))
	const funcName = "StringsSize16"

	for i := range fl {
		names[i] = fmt.Sprintf("%s.%s(%s.%s)", pkgName, funcName, fl[i].structTyp.receiver, fl[i].name)
	}
	declarations = declsZ(len(fl))
	declaration := strings.Join(declarations, ", ")

	return fmt.Sprintf("%s := %s",
		declaration,
		strings.Join(names, ", "),
	), declarations
}

func decls(u int, lenName string) (s []string) {
	for i := 0; i < u; i++ {
		s = append(s, lenName+strconv.Itoa(i))
	}
	return
}

func declsZ(u int) (s []string) {
	for i := 0; i < u; i++ {
		s = append(s, "z"+strconv.Itoa(i))
	}
	return
}
