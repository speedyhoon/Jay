package generate

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"strconv"
	"strings"
)

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

		var lv string
		if f.isArray() {
			lv = printFunc(
				nameOf(
					jay.SizeStringsArray,
					f.structTyp.isImportJ,
				),
				f.Field(""),
				strconv.Itoa(f.arraySize),
			)
		} else {
			lv = printFunc(
				nameOf(
					f.sizeOfPick(jay.SizeStrings8, jay.SizeStrings8),
					f.structTyp.isImportJ,
				),
				f.Name(),
			)
		}
		*values = append(*values, lv)

	} else {
		*values = append(*values, printFunc(lenKeyword, f.Name()))
	}

	*list = append(*list, string(f.marshal.qtyVar))
}
