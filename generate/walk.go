package generate

import (
	"go/ast"
)

type visitor struct {
	enclosing string
	structs   *[]structTyp
	option    Option
}

type structTyp struct {
	name                        string
	fixedLen, variableLen, bool fieldList // Exported fields.
}

type field struct {
	name string // The string used as the variable name.
	typ  string // The underlying type of the variable (uint, byte, bool, map, etc).

	// Alias name assigned to the type, for example, `type Toggle bool`, typ = "bool", aliasType = "Toggle".
	aliasType  string
	tag        string // The tag value within `j:""`
	tagOptions        // Valid tag options that have been successfully parsed and loaded from the `tag` string.
}
type fieldList []field

// Visit traverses the AST File and finds all structs even if they are unexported.
// Unexported structs can be exported if they are referenced in exported structs with exported field names.
// For example, type Cow struct { Id id }
func (v visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.TypeSpec:
		if n.Name != nil {
			v.enclosing = n.Name.Name
		}
	case *ast.StructType:
		if n.Fields == nil || len(n.Fields.List) == 0 {
			return v
		}

		s := structTyp{name: v.enclosing}
		if s.process(n.Fields.List, v.option) {
			*v.structs = append(*v.structs, s)
		}
	}

	return v
}
