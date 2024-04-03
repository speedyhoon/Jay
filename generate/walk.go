package generate

import (
	"go/ast"
)

type visitor struct {
	enclosing string
	structs   *[]structTyp
	option    Option
	dirList   *dirList
	dir       string
}

type field struct {
	name string // The string used as the variable name.
	typ  string // The underlying type of the variable (uint, byte, bool, map, etc).

	aliasType  string // Alias name assigned to the type, for example, `type Toggle bool`, typ = "bool", aliasType = "Toggle".
	arraySize  int    // 0 = not an array or slice, -1 = slice, >=1 = array size
	pkgReq     string // Which package is required to be imported if referenced in the generated code.
	arrayType  string // The type without the size in brackets. An empty string is not an array.
	tag        string // The tag value within `j:""`
	tagOptions        // Valid tag options that have been successfully parsed and loaded from the `tag` string.
	isFixedLen bool   // Is represented by a fixed quantity of bytes (like int64) or a variable quantity of bytes (like string & slices).
	isAliasDef bool
	isFirst    bool
	isLast     bool
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
		if n.Fields == nil || len(n.Fields.List) == 0 || v.enclosing == "" {
			return v
		}

		s := newStructTyp(v.dir, v.enclosing)
		if s.process(n.Fields.List, v.option, v.dirList) {
			*v.structs = append(*v.structs, s)
		}
	}

	return v
}
