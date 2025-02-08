package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/utl"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

const (
	StructTagName = "j" // StructTagName is the field tag to specify optional flags. For example: `j:max=0`
	IgnoreFlag    = "-" // IgnoreFlag is the value to ignore any exported field from serialising: `j:-`
	tagSymbol     = '`' // tagSymbol specifies which rune encloses field tags.
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

func isBuiltIn(typ string) bool {
	switch typ {
	case tBool, tByte, tFloat32, tFloat64, tInt, tInt8, tInt16, tInt32, tInt64, tString, tUint, tUint16, tUint32, tUint64:
		return true
	}
	return false
}

func (o Option) isSupportedType(t interface{}, dirList *dirList, pkg string) (f field, ok bool) {
	switch d := t.(type) {
	case *ast.Ident:
		if d.Obj == nil {
			typ := resolveBuiltinAlias(d.Name)
			if !isBuiltIn(typ) && dirList != nil {
				// Object might be declared in another file in the same package.
				d.Obj = findImportedType((*dirList)[pkg].files, (*dirList)[pkg].pkg, typ)
				if d.Obj == nil {
					return f, false
				}
				f, ok = o.isSupportedType(d.Obj, dirList, pkg)
				f.aliasType = typ
			}
			if f.typ == "" {
				f.typ = typ
			}
			if f.elmSize == 0 {
				f.elmSize = isLen(f.typ)
			}
			f.isFixedLen = o.isLenFixed(f.typ)
			return f, true
		}

		f, ok = o.isSupportedType(d.Obj, dirList, pkg)

	case nil:
	// Ignore.
	case *ast.SelectorExpr:
		f, ok = o.isSupportedSelector(d, dirList)

	case *ast.Object:
		if d.Kind != ast.Typ || d.Name == "" {
			lg.Println(d)
			return f, false
		}
		f, ok = o.isSupportedType(d.Decl, dirList, pkg)

	case *ast.ArrayType:
		f, ok = o.isSupportedType(d.Elt, dirList, pkg)
		if !ok {
			return f, false
		}
		f.arrayType = f.typ
		if f.isDef {
			// f.isDef prevents types like []float where `type float float32` which can't be easily converted to []float32 in one line.
			// However, `type float = float32`, `type floats = []float32` & `type floats []float32` can be easily converted in one line.
			if f.aliasType != tTimeDuration {
				return f, false
			}

			f.typ = "[]" + f.aliasType
		} else {
			f.typ = "[]" + f.typ
		}

		f.arraySize, ok = calcArraySize(d.Len)
		f.isFixedLen = f.isFixedLen && f.isArray()

	case *ast.TypeSpec:
		f, ok = o.isSupportedType(d.Type, dirList, pkg)
		// Field is a type definition if isDef has already been set via inspecting d.Type beforehand, like time.Duration (type Duration int64),
		// OR if TypeSpec is not an assignment (there is no equal sign).
		f.isDef = f.isDef || d.Assign == token.NoPos

		// Saves having to run goimports to fix unused imports.
		if f.aliasType != d.Name.Name {
			// Replace the imported type with the local type name.
			f.aliasType = d.Name.Name

			if f.isDef && isBuiltIn(f.typ) {
				// Remove the import required when the underlying type is a simple
				// built-in type (like int64) instead of an imported struct (like time.Time)
				// that needs the import, to resolve any type conversions needed.
				// This is only needed for type definitions (type A int), not type aliases (type A = int).
				f.pkgReq = ""
			}
		}

	default:
		lg.Printf("type %T not expected in Option.isSupportedType()", d)
	}
	return
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

// isSupportedSelector resolves imported types and some types within Go's standard library.
func (o Option) isSupportedSelector(d *ast.SelectorExpr, dirList *dirList) (f field, ok bool) {
	x, ok := d.X.(*ast.Ident)
	if !ok {
		return
	}

	// Go standard library types are hardcoded rather than trying to find the Go source files on every environment;
	// especially difficult when the GOROOT environment variable isn't set and Go could be installed anywhere, even a NAS.
	switch x.Name {
	case "time":
		switch d.Sel.Name {
		case "Duration": // type Duration int64
			f.typ = tInt64
			f.pkgReq = x.Name
			f.aliasType = tTimeDuration
			f.isDef = true
			f.isFixedLen = true
			f.elmSize = isLen(f.typ)
			return f, true
		case "Time": // type Time struct
			f.typ = tTime
			f.pkgReq = x.Name
			f.aliasType = tTime
			f.isFixedLen = true
			return f, true
		}
	}

	obj := findImportedType(dirList.allFiles(), x.Name, d.Sel.Name)
	if obj != nil {
		f, ok = o.isSupportedType(obj, nil, "")
		f.aliasType = d.Sel.Name
		return f, f.typ != ""
	}

	return
}

func pkgSelName(pkg, selector string) string {
	if pkg != "" {
		return fmt.Sprintf("%s.%s", pkg, selector)
	}
	return selector
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
func genType(arraySize int, typ string) string {
	switch arraySize {
	case typeSlice:
		return fmt.Sprintf("[]%s", typ)
	case typeNotArrayOrSlice:
		return typ
	default:
		return fmt.Sprintf("[%d]%s", arraySize, typ)
	}
}

const (
	typeSlice           = -1
	typeNotArrayOrSlice = 0
	typeArray           = 1
)

// calcArraySize returns -1 for a slice, 0 as invalid & >= 1 for array size.
func calcArraySize(x interface{}) (size int, ok bool) {
	switch d := x.(type) {
	case nil:
		return typeSlice, true
	case *ast.BasicLit:
		if d.Kind != token.INT {
			lg.Println("unhandled token type", d.Kind, d.Value)
			return typeNotArrayOrSlice, false
		}

		u, err := strconv.ParseUint(d.Value, 10, 64)
		if err != nil {
			lg.Println("invalid array size", d.Value)
			return typeNotArrayOrSlice, false
		}
		if u > maxUint24 {
			lg.Println("invalid array size", d.Value, "> expected:", maxUint24)
			return typeNotArrayOrSlice, false
		}
		if u == 0 {
			lg.Println("arrays with size zero ([0]int) are not supported")
			return typeNotArrayOrSlice, false
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
	return typeNotArrayOrSlice, false
}

func (s *structTyp) hasExportedFields() bool {
	return len(s.fixedLen) >= 1 || len(s.variableLen) >= 1 || len(s.bool) >= 1 || len(s.single) >= 1 || len(s.stringSlice) >= 1
}

func (s *structTyp) addExportedFields(names []*ast.Ident, f field) {
	f.structTyp = s
	for m := range names {
		f.name = names[m].Name
		if f.typ == tBool || f.arrayType == tBool && f.isArray() {
			s.bool = append(s.bool, f)
			continue
		}
		switch f.typ {
		case tByte, tInt8:
			s.single = append(s.single, f)
			continue
		case tStrings:
			s.stringSlice = append(s.stringSlice, f)
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

// isLen returns how many bytes each type requires.
func isLen(typ string) uint {
	switch typ {
	case tBool, tByte, tInt8, tString:
		return 1
	case tInt16, tUint16:
		return 2
	case tInt32, tUint32, tFloat32:
		return 4
	case tInt64, tUint64, tFloat64:
		return 8
	case tInt, tUint:
		// TODO dynamically set size for 32 or 64 bit systems
		return 8
	}
	return 0
}

func (o Option) isLenFixed(typ string) bool {
	switch typ {
	case tInt:
		return o.FixedIntSize
	case tString:
		return false
	case tUint:
		return o.FixedUintSize
	}
	return true
}

func bufWriteF(b *bytes.Buffer, format string, a ...any) {
	_, err := fmt.Fprintf(b, format, a...)
	if err != nil {
		lg.Println("err writing to buffer:", err)
	}
}

func (s *structTyp) defineTrackingVars(buf *bytes.Buffer, byteIndex uint) (at, end string) {
	switch len(s.variableLen) {
	case 0:
		return
	case 1:
		at = utl.UtoA(byteIndex)
	default:
		if s.variableLen[0].typ == tBoolS {
			bufWriteF(buf, "at, end := %d, %[1]d+%s(%s)\n", byteIndex, nameOf(jay.SizeBools, nil), lenVariable(0))
		} else {
			bufWriteF(buf, "at, end := %d, %[1]d+%s\n", byteIndex, multiples(s.variableLen[0], lenVariable(0)))
		}
		at, end = "at", "end"
	}
	return
}

func (s *structTyp) tracking(buf *bytes.Buffer, i int, endVar string, byteIndex uint, varType string) (at, end string) {
	if endVar == "" {
		return utl.UtoA(byteIndex), ""
	}

	if i == len(s.variableLen)-1 {
		return "end", ""
	}
	if i >= 1 {
		if varType == tBoolS {
			bufWriteF(buf, "at, end = end, end+%s(%s)\n", nameOf(jay.SizeBools, nil), lenVariable(i))
		} else {
			bufWriteF(buf, "at, end = end, end+%s\n", multiples(s.variableLen[i], lenVariable(i)))
		}
	}
	return "at", "end"
}

func lenVariable(index int) string {
	return lengthNameDefault + utl.UtoA(uint(index))
}

// List of supported types.
const (
	tBool, tBoolS                            = "bool", "[]bool"
	tByte, tByteS                            = "byte", "[]byte"
	tFloat32, tFloat32S                      = "float32", "[]float32"
	tFloat64, tFloat64S                      = "float64", "[]float64"
	tInt, tInt8, tInt16, tInt32, tInt64      = "int", "int8", "int16", "int32", "int64"
	tIntS, tInt8S, tInt16S, tInt32S, tInt64S = "[]int", "[]int8", "[]int16", "[]int32", "[]int64"
	tUint, tUint16, tUint32, tUint64         = "uint", "uint16", "uint32", "uint64"
	tUintS, tUint16S, tUint32S, tUint64S     = "[]uint", "[]uint16", "[]uint32", "[]uint64"
	tString, tStrings                        = "string", "[]string"
	tTime, tTimeDuration, tTimeDurations     = "time.Time", "time.Duration", "[]time.Duration"
)

// resolveBuiltinAlias replaces the built-in alias with the underlining name to reduce the quantity of types to support.
func resolveBuiltinAlias(typ string) string {
	switch typ {
	case "uint8":
		return tByte
	case "rune":
		return tInt32
	case tTimeDuration:
		return tInt64
	}
	return typ
}
