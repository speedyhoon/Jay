package jay

import (
	"go/ast"
	//"github.com/dave/ast"
	//"github.com/dave/ast/decorator"
)

type visitor struct {
	enclosing string
	structs   *[]Struct
}

type Struct struct {
	name                  string
	fields                []*ast.Field
	fixedLen, variableLen []field // Exported fields.
}

type field struct {
	name, tag, typ, typName string
	Options
}

// Visit traverses the AST File.
func (v visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.TypeSpec:
		if n.Name != nil && ast.IsExported(n.Name.Name) {
			v.enclosing = n.Name.Name
		}
	case *ast.StructType:
		if n.Fields != nil && len(n.Fields.List) != 0 && ast.IsExported(v.enclosing) {
			*v.structs = append(*v.structs, Struct{name: v.enclosing, fields: n.Fields.List})
		}
	}
	return v
}
