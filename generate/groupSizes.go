package generate

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"sort"
	"strings"
)

type varSize map[uint][]string

func (mm *varSize) add(sizeOf uint, varName string) {
	if sizeOf == 0 {
		varName = printFunc(nameOf(jay.SizeBools, nil), varName)
	}

	_, ok := (*mm)[sizeOf]
	if !ok {
		(*mm)[sizeOf] = []string{varName}
		return
	}

	(*mm)[sizeOf] = append((*mm)[sizeOf], varName)
}

func (mm *varSize) group() (output string) {
	c, l := 0, len(*mm)
	sizes := make([]uint, len(*mm))
	for u := range *mm {
		sizes[c] = u
		c++
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	var grouped []string
	for i := 0; i < l; i++ {
		sizeOf := sizes[i]
		additionJoin((*mm)[sizeOf], sizeOf, &grouped)
	}
	return strings.Join(grouped, "+")
}

func additionJoin(list []string, sizeOf uint, out *[]string) {
	if len(list) >= 1 {
		*out = append(*out, formatVarAdditions(list, sizeOf))
	}
}

func formatVarAdditions(list []string, sizeOf uint) string {
	if len(list) == 1 {
		if sizeOf <= 1 { // bools and singles.
			return list[0]
		}
		return fmt.Sprintf("%d*%s", sizeOf, list[0])
	}
	if sizeOf <= 1 { // bools and singles.
		return strings.Join(list, "+")
	}
	return fmt.Sprintf("%d*(%s)", sizeOf, strings.Join(list, "+"))
} //67
