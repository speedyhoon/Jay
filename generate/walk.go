package generate

import (
	"fmt"
	"go/ast"
)

type visitor struct {
	enclosing string
	structs   *[]*structTyp
	option    *Option
	dirList   *dirList
	dir       string
}

type field struct {
	name string // The string used as the variable name.
	typ  string // The underlying type of the variable (uint, byte, bool, map, etc).

	structTyp *structTyp // Pointer to the parent struct containing this field.
	fieldList *fieldList

	aliasType  string // Alias name assigned to the type. For example, `type Toggle bool`, field.typ = "bool", field.aliasType = "Toggle".
	arraySize  int    // 0 = not an array or slice, -1 = slice, >=1 = array size.
	elmSize    uint   // Quantity of bytes required to represent the type. For example, bool = 0, int8 = 1, uint16 = 2, int64 = 8, etc.
	pkgReq     string // The third party package required to be imported if referenced in the generated code.
	arrayType  string // The type without the size in brackets. An empty string is not an array.
	lenVar     string // For variable length types, this is the name of the variable that stores the length from len(b[X]) in marshal methods.
	marshal    vars
	unmarshal  vars
	qtyIndex   []uint // Stores which byte indexes to read/write the qty as uint8, uint16 or uint24 for variable length types.
	indexStart *uint  // Byte index to start at. Variable length fields might be separated from their qtyIndex.
	indexEnd   *uint  // Byte index to end at if type is fixed length.
	tag        string // The tag value within `j:""`
	tagOptions        // Valid tag options that have been successfully parsed and loaded from the `tag` string.
	isFixedLen bool   // Is represented by a fixed quantity of bytes (like int64) or a variable quantity of bytes (like string and slices).
	isDef      bool   // Whether the type is a new definition (type num int) verses an alias (type num = int) or a normal type (int).
	isFirst    bool   // True when this field is the first to be marshalled and unmarshalled in the struct methods (not the first defined in the struct definition).
	isLast     bool
}
type fieldList []*field

// Visit traverses the AST File and finds all structs even if they are unexported.
// Unexported structs can be exported if they are referenced in exported structs with exported field names.
// For example, type Cow struct { ID id }
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

		s := newStructTyp(v.dir, v.enclosing, v.option)
		if s.process(n.Fields.List, v.dirList) {
			*v.structs = append(*v.structs, s)
		}
	}

	return v
}

func (f *field) Name() string {
	return fmt.Sprintf("%s.%s", f.structTyp.receiver, f.name)
}
