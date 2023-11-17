package jay

import (
	"go/ast"
)

type visitor struct {
	enclosing string
	structs   *[]Struct
}

type Struct struct {
	name                  string
	fixedLen, variableLen []field // Exported fields.
}

type field struct {
	name       string // The string used as the variable name.
	typ        string // The type of the variable (uint, byte, bool, etc).
	typName    string // Alias name assigned to the type, for example, `var Toggle bool`.
	tag        string // The tag value within `j:""`
	tagOptions        // Valid tag options that have been successfully parsed and loaded from the `tag` string.
}

type tagOptions struct {
	// The maximum and minimum value expected in the variable.
	// Any value out of this range isn't guaranteed to be marshaled or unmarshaled correctly.
	Max, Min uint

	maxBytes uint
}

// Visit traverses the AST File.
func (v visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.TypeSpec:
		if n.Name != nil && ast.IsExported(n.Name.Name) {
			v.enclosing = n.Name.Name
		}
	case *ast.StructType:
		if n.Fields == nil || len(n.Fields.List) == 0 || !ast.IsExported(v.enclosing) {
			return v
		}

		s := Struct{name: v.enclosing}
		if s.Process(n.Fields.List) {
			*v.structs = append(*v.structs, s)
		}
	}

	return v
}
