package generate

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"strings"
)

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
				"%s.%s(%s.%s)",
				pkgName,
				nameOf(jay.StringsSize16, nil),
				receiver,
				slices[i].name,
			))
		}
		out = strings.Join(sizes, ", ")
		if len(names) == 0 {
			declarations := strings.Join(decls(qty), ", ")
			return fmt.Sprintf("%s := %s",
				declarations,
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
		s = append(s, lenVariable(i))
	}
	return
}
