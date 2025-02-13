package generate

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"strings"
)

func multiples(f *field) (lenVar string) {
	switch {
	case f.isSlice():
		itemSize := field{typ: f.arrayType, structTyp: f.structTyp}.typeFuncSize()
		if itemSize >= 2 {
			return fmt.Sprintf("%s*%d", f.lenVar, itemSize)
		}
		return f.lenVar
	case f.isArray():
		itemSize := field{typ: f.arrayType, structTyp: f.structTyp}.typeFuncSize()
		return utl.UtoA(uint(f.arraySize) * itemSize)
	default: // typeNotArrayOrSlice
		return f.lenVar
	}
}

func (s *structTyp) varLenFieldNames() (names []string) {
	for _, v := range s.variableLen {
		names = append(names, v.Name())
	}
	return
}

func (s *structTyp) generateLenVarLine() string {
	var names, values []string
	for i, f := range append(s.stringSlice, s.variableLen...) {
		f.generateLenVar(&names, &values, i)
	}

	if len(names) == 0 {
		return ""
	}

	return fmt.Sprintf(
		"%s:=%s",
		strings.Join(names, ","),
		strings.Join(values, ","),
	)
}

func (f *field) generateLenVar(list, values *[]string, index int) {
	if f.fieldList == &f.structTyp.stringSlice {
		lv := printFunc(
			nameOf(
				f.sizeOfPick(jay.StringsSize8, jay.StringsSize16),
				nil,
			),
			f.Name(),
		)

		if f.isLast {
			f.lenVar = lv
			return
		}

		*values = append(*values, lv)

	} else {
		*values = append(*values, printFunc("len", f.Name()))
	}

	f.lenVar = lenVariable(index)
	*list = append(*list, f.lenVar)
}
