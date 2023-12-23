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
	name       string
	receiver   string
	bufferName string

	// Exported fields.
	fixedLen, // Fixed length types like int16, uint64 and some arrays etc.
	single, // Fields represented in one byte like int8 & uint8. These have additional optimisations.
	variableLen, // Variable length fields like string and all slice types. These are generated last & have the most processing overhead.
	bool fieldList // Boolean fields are joined together and represented as binary.
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

		s := structTyp{
			name:     v.enclosing,
			receiver: receiverName(v.enclosing),
		}
		s.bufferName = bufferName(s.receiver)

		if s.process(n.Fields.List, v.option) {
			*v.structs = append(*v.structs, s)
		}
	}

	return v
}
