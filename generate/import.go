package generate

import (
	"fmt"
	"strings"
)

type importList []string

func (m *importList) add(pkg string) {
	if pkg == "" {
		return
	}

	// Check every import doesn't already exist in the list.
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

func (m importList) Print() string {
	switch len(m) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("import \"%s\"\n", m[0])
	default:
		return fmt.Sprintf("import (\n\t\"%s\"\n)\n", strings.Join(m, "\"\n\t\""))
	}
}
