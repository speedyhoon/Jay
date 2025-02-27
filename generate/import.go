package generate

import (
	"fmt"
	"sort"
	"strings"
)

type importList []string

func (m *importList) add(pkg string) {
	if pkg == "" {
		return
	}

	// Check if the new import doesn't already exist in the list.
	for i := range *m {
		if (*m)[i] == pkg {
			return
		}
	}

	*m = append(*m, pkg)
}

func (m *importList) join(l importList) {
	for i := range l {
		m.add(l[i])
	}
}

func (m importList) print() string {
	switch len(m) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("\nimport \"%s\"\n", m[0])
	default:
		// Sort imports by their package name.
		sort.Strings(m)
		return fmt.Sprintf("\nimport (\n\t\"%s\"\n)\n", strings.Join(m, "\"\n\t\""))
	}
}
