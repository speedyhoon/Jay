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

func lengths2(names []string, slices fieldList, receiver string) string {
	if len(names) == 0 && len(slices) <= 1 {
		return ""
	}

	qty := len(names)
	var sizes []string
	var out string
	if l := len(slices) - 1; l >= 1 {
		qty += l

		for i := 0; i < l; i++ {
			sizes = append(sizes, fmt.Sprintf(
				"jay.StringsSize16(%s.%s)",
				receiver,
				slices[i].name,
			))
		}
		out = strings.Join(sizes, ", ")
		if len(names) == 0 {
			return fmt.Sprintf("%s := %s",
				strings.Join(decls(qty), ", "),
				out,
			)
		}

		out = ", " + out
	}

	receiver = fmt.Sprintf("len(%s.", receiver)
	declarations := strings.Join(decls(qty), ", ")
	return fmt.Sprintf("%s := %s%s)%s",
		declarations,
		receiver,
		strings.Join(names, "),"+receiver),
		out,
	)
}

func decls(u int) (s []string) {
	for i := 0; i < u; i++ {
		s = append(s, "l"+strconv.Itoa(i))
	}
	return
}
