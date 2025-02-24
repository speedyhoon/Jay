package generate

import (
	"fmt"
	"sort"
	"strings"
)

type varSize map[uint][]string

func (vs *varSize) add(f *field, varName string) {
	sizeOf := f.elmSize

	_, ok := (*vs)[sizeOf]
	if !ok {
		(*vs)[sizeOf] = []string{varName}
		return
	}

	(*vs)[sizeOf] = append((*vs)[sizeOf], varName)
}

func (vs *varSize) group() (output string) {
	c, l := 0, len(*vs)
	sizes := make([]uint, len(*vs))
	for u := range *vs {
		sizes[c] = u
		c++
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	var grouped []string
	for i := 0; i < l; i++ {
		sizeOf := sizes[i]
		additionJoin((*vs)[sizeOf], sizeOf, &grouped)
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
}
