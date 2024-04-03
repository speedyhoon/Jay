package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"go/ast"
	"go/token"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	StructTagName = "j"
	IgnoreFlag    = "-"
	tagSymbol     = '`'
)

// getTag returns the value associated with key "j" in the tag string.
func getTag(b *ast.BasicLit) string {
	if b == nil {
		return ""
	}

	if b.Kind != token.STRING {
		return ""
	}

	return strings.TrimSpace(reflect.StructTag(unwrapTagValue(b.Value)).Get(StructTagName))
}

// unwrapTagValue removes the leading and trailing grave (`) if present.
func unwrapTagValue(str string) string {
	if len(str) >= 2 && str[0] == tagSymbol && str[len(str)-1] == tagSymbol {
		return str[1 : len(str)-1]
	}
	return str
}

var typeArray = regexp.MustCompile(`\[(\d)+\](\w)+`)
var typeArrayBrackets = regexp.MustCompile(`\[(\d)+\]`)

func supportedType(typ string) bool {
	switch typ {
	case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "[]byte":
		return true
	}

	if typeArray.MatchString(typ) {
		switch typeArrayBrackets.ReplaceAllString(typ, "") {
		case "byte", "int8", "uint8":
			return true
		}
	}

	return false
}

func (o Option) isSupportedType(t *ast.Field, dirList *dirList, pkg string) (f field, ok bool) {
	switch d := t.Type.(type) {
	case *ast.Ident:
		if supportedType(d.Name) { // TODO its probably better to check if the obj != nil first otherwise do this.
			f.typ = d.Name
			f.aliasType = d.Name
			f.isFixedLen = o.isLenFixed(d.Name)
			return f, true
		}

		// Object might be declared in another file in the same package.
		if d.Obj == nil && dirList != nil {
			d.Obj = findImportedType((*dirList)[pkg].files, pkg, t.Names[0].Name)
		}

		f = o.typeOf(d.Obj)
		if f.isAliasDef {
			// type definition (`type toggle bool`) requires converting (`bool(foo)`) or casting (`foo.(bool)`) to correct type.
			f.aliasType = d.Name
		} else {
			// type alias (`type toggle = bool`) can be used interchangeably. No extra processing required.
			f.aliasType = f.typ
		}
		return f, f.typ != ""
	case nil:
	// Ignore.
	case *ast.SelectorExpr:
		if dirList != nil {
			return o.isSupportedSelector(d, dirList.allFiles())
		} else {
			return o.isSupportedSelector(d, nil)
		}

	//case *ast.StructType:
	// TODO not yet implemented.
	case *ast.ArrayType:
		f, ok = o.calcType(d, "", dirList.allFiles())
		return f, ok
	default:
		lg.Printf("type %T not expected in Option.isSupportedType()", d)
	}
	return f, false
}

func findImportedType(files []*ast.File, pkg, typName string) *ast.Object {
	// TODO doesn't handle aliased imports
	var found []*ast.Object

	for _, file := range files {
		if file == nil || file.Name == nil || file.Name.Name != pkg || file.Scope == nil {
			continue
		}

		for key, obj := range file.Scope.Objects {
			if obj == nil || obj.Decl == nil {
				continue
			}
			if key == typName {
				found = append(found, obj)
				//return obj
			}
		}
	}

	switch len(found) {
	case 0:
		lg.Println("no imports found:", pkgSelName(pkg, typName))
	case 1:
		return found[0]
	default:
		lg.Printf("%d imports found: %s", len(found), pkgSelName(pkg, typName))
	}
	return nil
}

func (o Option) calcType(t interface{}, typePrefix string, files []*ast.File) (f field, ok bool) {
	switch d := t.(type) {
	case *ast.Field:
		return o.calcType(d.Type, typePrefix, files)
	case *ast.Ident:
		name := typePrefix + d.Name
		if supportedType(name) {
			f.typ = name
			f.aliasType = name
			f.isFixedLen = o.isLenFixed(name)
			return f, true
		}

		// Type has an alias name.
		f = o.typeOf(d.Obj)
		f.aliasType = name
		return f, f.typ != ""
	case nil:
	// Ignore.
	case *ast.SelectorExpr:
		return o.isSupportedSelector(d, files)
	//case *ast.StructType:
	// TODO not yet implemented.
	case *ast.ArrayType:
		var typ string
		typ, ok = o.calcArrayType(d.Elt)
		if !ok {
			return f, false
		}
		var size int
		size, ok = calcArraySize(d.Len)
		if !ok {
			return f, false
		}
		f = o.newFieldArray(size, typ)
		return f, true

	case *ast.BasicLit:
		if d.Kind == token.INT {
			return f, true
		}
	default:
		lg.Printf("type %T not expected in Option.isSupportedType()", d)
	}
	return f, false
}

func (o Option) isSupportedSelector(d *ast.SelectorExpr, files []*ast.File) (f field, ok bool) {
	x, ok := d.X.(*ast.Ident)
	if !ok {
		return
	}

	switch x.Name {
	case "time":
		switch d.Sel.Name {
		case "Time", "Duration":
			f.typ = pkgSelName(x.Name, d.Sel.Name)
			f.aliasType = f.typ
			f.pkgReq = x.Name
			f.isFixedLen = true
			return f, true
		}
	}

	obj := findImportedType(files, x.Name, d.Sel.Name)
	if obj != nil {
		f = o.typeOf(obj)
		f.aliasType = d.Sel.Name
		return f, f.typ != ""
	}

	return
}

func pkgSelName(pkg, selector string) string {
	return fmt.Sprintf("%s.%s", pkg, selector)
}

func packageName(f *ast.File) string {
	if f != nil && f.Name != nil {
		return f.Name.Name
	}
	return "main"
}

func (o *Option) newFieldArray(arraySize int, arrayType string) (f field) {
	f = field{
		arraySize:  arraySize,
		arrayType:  arrayType,
		typ:        genType(arraySize, arrayType),
		isFixedLen: arraySize >= 0 && o.isLenFixed(arrayType),
	}
	return
}
func genType(arraySize int, arrayType string) string {
	switch {
	case arraySize <= -1:
		return fmt.Sprintf("[]%s", arrayType)
	case arraySize == 0:
		return arrayType
	default:
		return fmt.Sprintf("[%d]%s", arraySize, arrayType)
	}
}

func (o *Option) calcArrayType(x interface{}) (typeOf string, ok bool) {
	switch d := x.(type) {
	case *ast.Ident:
		if supportedArrayType(d.Name) {
			return d.Name, true
		}
	}
	lg.Println("array type not supported yet")
	return "", false
}
func supportedArrayType(s string) bool {
	switch s {
	case "byte", "int8", "uint8", "bool":
		return true
	default:
		lg.Println("array type", s, "not supported yet")
	}
	return false
}

// calcArraySize returns -1 for a slice, 0 as invalid & >= 1 for array size.
func calcArraySize(x interface{}) (size int, ok bool) {
	switch d := x.(type) {
	case nil:
		return -1, true
	case *ast.BasicLit:
		if d.Kind != token.INT {
			lg.Println("unhandled token type", d.Kind, d.Value)
			return 0, false
		}

		u, err := strconv.ParseUint(d.Value, 10, 24)
		if err != nil {
			lg.Println("invalid array size", d.Value)
			return 0, false
		}
		return int(u), true

	case *ast.ValueSpec:
		if len(d.Values) == 1 {
			return calcArraySize(d.Values[0])
		}
	case *ast.Ident:
		if d.Obj != nil {
			return calcArraySize(d.Obj.Decl)
		}
	default:
		lg.Printf("unhandled case %t in calcArraySize", d)
	}
	return 0, false
}

func (o Option) typeOf(t interface{}) (fe field) {
	switch x := t.(type) {
	case *ast.Object:
		if x == nil || x.Name == "" || x.Kind != ast.Typ {
			return
		}
		return o.typeOf(x.Decl)
	case *ast.TypeSpec:
		fe = o.typeOf(x.Type)
		fe.isAliasDef = fe.isAliasDef || x.Assign == token.NoPos
		return
	case *ast.StructType:
		if x.Fields != nil && len(x.Fields.List) != 0 {
			return field{typ: "struct", isFixedLen: !o.isVariableLen(x.Fields.List)}
		}
	case *ast.Ident:
		if x.Obj != nil {
			return o.typeOf(x.Obj)
		}

		if supportedType(x.Name) {
			return field{typ: x.Name, isFixedLen: o.isLenFixed(x.Name)}
		}
	case *ast.SelectorExpr:
		fe, _ = o.isSupportedSelector(x, nil)
		return
	case nil:
		// Ignore.
	default:
		lg.Printf("type %T not expected in typeOf", x)
	}
	return
}

func (o Option) isVariableLen(fields []*ast.Field) bool {
	for _, f := range fields {
		if !hasExported(f.Names) {
			continue
		}

		fe, ok := o.isSupportedType(f, nil, "")
		if !ok {
			continue
		}

		if !fe.isFixedLen {
			return true
		}
	}
	return false
}

func (s *structTyp) hasExportedFields() bool {
	return len(s.fixedLen) >= 1 || len(s.variableLen) >= 1 || len(s.bool) >= 1 || len(s.single) >= 1
}

func hasExported(idents []*ast.Ident) bool {
	for _, ident := range idents {
		if ident.IsExported() {
			return true
		}
	}
	return false
}

func (s *structTyp) addExportedFields(names []*ast.Ident, f field) {
	for m := range names {
		f.name = names[m].Name
		switch f.typ {
		case "bool":
			s.bool = append(s.bool, f)
			continue
		case "byte", "int8", "uint8":
			s.single = append(s.single, f)
			continue
		}
		// TODO add support for adding tiny enums using <= 7 bits

		if f.isFixedLen {
			s.fixedLen = append(s.fixedLen, f)
		} else {
			s.variableLen = append(s.variableLen, f)
		}
	}
}

//func isLenFixed(typ string) bool {
//	switch typ {
//	case "byte" /*"float32","float64",*/, "int", "int8", "int16", "int32", "int64", "rune", "uint", "uint8", "uint16", "uint32", "uint64":
//		return true
//	}
//	return false
//}

func isLen(typ string) uint {
	switch typ {
	case "bool", "byte", "int8", "uint8", "string", "int", "uint":
		return 1
	case "int16", "uint16":
		return 2
	case "int32", "rune", "uint32", "float32":
		return 4
	case "int64", "uint64", "float64":
		return 8
	}
	return 0
}

func (o Option) isLenFixed(typ string) bool {
	switch typ {
	case "int":
		return o.FixedIntSize
	case "string", "[]byte", "[]uint8", "[]int8", "[]bool":
		return false
	case "uint":
		return o.FixedUintSize
	}
	return true
}

func (o Option) isLenVariable(typ string) bool {
	return !o.isLenFixed(typ)
}

/*func isLenStructVariable(t interface{}) bool {
	switch x := t.(type) {
	case *ast.Ident:
		if x.Obj != nil {
			return isLenStructVariable(x.Obj.Decl)
		}
	case nil:
	case *ast.TypeSpec:
		if x.Type != nil {
			return isLenStructVariable(x.Type)
		}
	case *ast.StructType:

		return isLenStructVariable(x.Fields.List)
	}
	return false
}*/

func bufWriteF(b *bytes.Buffer, format string, a ...any) {
	b.WriteString(fmt.Sprintf(format, a...))
}

func (s *structTyp) defineTrackingVars(buf *bytes.Buffer, byteIndex uint) (at, end string) {
	switch len(s.variableLen) {
	case 0:
		return
	case 1:
		at = Utoa(byteIndex)
	default:
		if s.variableLen[0].typ == "[]bool" {
			bufWriteF(buf, "at, end := %d, %[1]d+%s(%s)\n", byteIndex, nameOf(jay.SizeBools, nil), lenVariable(0))
		} else {
			bufWriteF(buf, "at, end := %d, %[1]d+%s\n", byteIndex, lenVariable(0))
		}
		at, end = "at", "end"
	}
	return
}

func (s *structTyp) tracking(buf *bytes.Buffer, i int, endVar string, byteIndex uint, varType string) (at, end string) {
	if endVar == "" {
		return Utoa(byteIndex), ""
	}

	if i == len(s.variableLen)-1 {
		return "end", ""
	}
	if i >= 1 {
		if varType == "[]bool" {
			bufWriteF(buf, "at, end = end, end+%s(%s)\n", nameOf(jay.SizeBools, nil), lenVariable(i))
		} else {
			bufWriteF(buf, "at, end = end, end+%s\n", lenVariable(i))
		}
	}
	return "at", "end"
}

func lenVariable(index int) string {
	const lengthVarPrefix = "l"
	return lengthVarPrefix + Utoa(uint(index))
}
