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
		fe := field{typ: f.arrayType, structTyp: f.structTyp}
		itemSize := fe.typeFuncSize()
		if itemSize >= 2 {
			return fmt.Sprintf("%s*%d", f.marshal.qtyVar, itemSize)
		}
		return string(f.marshal.qtyVar)
	case f.isArray():
		fe := field{typ: f.arrayType, structTyp: f.structTyp}
		itemSize := fe.typeFuncSize()
		return utl.UtoA(uint(f.arraySize) * itemSize)
	default: // typeNotArrayOrSlice
		return string(f.marshal.qtyVar)
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
	for _, f := range append(s.stringSlice, s.variableLen...) {
		f.generateLenVar(&names, &values)
	}

	if len(names) == 0 {
		return ""
	}

	return fmt.Sprintf(
		"%s := %s",
		strings.Join(names, ", "),
		strings.Join(values, ", "),
	)
}

func (f *field) generateLenVar(list, values *[]string) {
	if f.fieldList == &f.structTyp.stringSlice {
		if f.isLast {
			return
		}

		lv := printFunc(
			nameOf(
				f.sizeOfPick(jay.StringsSize8, jay.StringsSize16),
				nil,
			),
			f.Name(),
		)
		*values = append(*values, lv)

	} else {
		*values = append(*values, printFunc(lenKeyword, f.Name()))
	}

	*list = append(*list, string(f.marshal.qtyVar))
}
