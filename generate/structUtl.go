package generate

import (
	"bytes"
	"fmt"
	"go/token"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/dave/dst"
	"github.com/speedyhoon/utl"
)

const (
	StructTagName = "j" // StructTagName is the field tag to specify optional flags. For example, `j:max=0`
	IgnoreFlag    = "-" // IgnoreFlag is the value to ignore any exported field from serializing: `j:-`
	tagSymbol     = '`' // tagSymbol specifies which rune encloses field tags.
)

// getTag returns the value associated with key "j" in the tag string.
func getTag(b *dst.BasicLit) string {
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
	case tBool, tByte, tFloat32, tFloat64, tInt, tInt8, tInt16, tInt32, tInt64, tString, tUint, tUint16, tUint32, tUint64, tComplex64, tComplex128:
		return true
	}
	return false
}

func (s *structTyp) isSupportedType(f *field, t interface{}, dirList *dirList, pkg string, fileImports []*dst.ImportSpec, parents ...[]*dst.Ident) (ok bool) {
	switch d := t.(type) {
	case *dst.Ident:
		if d.Obj == nil {
			typ := resolveBuiltinAlias(d.Name)
			if !isBuiltIn(typ) && dirList != nil {
				// Object might be declared in another file in the same package.
				d.Obj = findImportedType((*dirList)[pkg].files, (*dirList)[pkg].pkg, typ)
				if d.Obj == nil {
					return false
				}
				ok = s.isSupportedType(f, d.Obj, dirList, pkg, fileImports, parents...)
				if !ok {
					return false
				}
				f.aliasType = typ
			}
			if f.typ == "" {
				f.typ = typ
			}
			if f.elmSize == 0 {
				f.elmSize = s.option.isLen(f.typ)
			}
			f.isFixedLen = s.option.isLenFixed(f.typ)
			return true
		}

		ok = s.isSupportedType(f, d.Obj, dirList, pkg, fileImports, parents...)

	// Ignore.
	case nil:

	case *dst.SelectorExpr:
		ok = s.option.isSupportedSelector(f, d, fileImports)

	case *dst.Object:
		if d.Kind != dst.Typ || d.Name == "" {
			lg.Println(d)
			return false
		}
		ok = s.isSupportedType(f, d.Decl, dirList, pkg, fileImports, parents...)

	case *dst.ArrayType:
		ok = s.isSupportedType(f, d.Elt, dirList, pkg, fileImports, parents...)
		if !ok {
			return false
		}
		f.arrayType = f.typ
		if f.isDef {
			// f.isDef prevents types like []float where `type float float32` which can't be easily converted to []float32 in one line.
			// However, `type float = float32`, `type floats = []float32` & `type floats []float32` can be easily converted in one line.
			if f.aliasType != tTimeDuration {
				return false
			}

			f.typ = "[]" + f.aliasType
		} else {
			f.typ = "[]" + f.typ
		}

		f.arraySize, ok = calcArraySize(d.Len)
		f.unmarshal.qtyVar = multiplier(strconv.Itoa(f.arraySize))
		f.marshal.qtyVar = multiplier(strconv.Itoa(f.arraySize))
		f.isFixedLen = f.isFixedLen && f.isArray()

	case *dst.TypeSpec:
		ok = s.isSupportedType(f, d.Type, dirList, pkg, fileImports, parents...)
		// Field is a type definition if isDef has already been set via inspecting d.Type beforehand, like time.Duration (type Duration int64),
		// OR if TypeSpec is not an assignment (there is no equal sign).
		f.isDef = f.isDef || !d.Assign

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

	case *dst.StructType:
		if d.Fields == nil || len(d.Fields.List) < 1 {
			return false
		}

		s.process(d.Fields.List, dirList, fileImports, parents...)

	default:
		lg.Printf("type %T not expected in Option.isSupportedType()", d)
	}
	return
}

func findImportedType(files []*dst.File, pkg, typName string) *dst.Object {
	// TODO doesn't handle aliased imports
	var found []*dst.Object

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
func (o Option) isSupportedSelector(f *field, d *dst.SelectorExpr, fileImports []*dst.ImportSpec) (ok bool) {
	x, ok := d.X.(*dst.Ident)
	if !ok {
		return
	}

	// Go standard library types are hardcoded rather than trying to find the Go source files on every environment;
	// especially difficult when the GOROOT environment variable isn't set and Go could be installed anywhere, even a NAS.
	switch x.Name {
	case "time":
		// Built-ins are hardcoded here so resolveImportedTypes doesn't need to be executed.
		switch d.Sel.Name {
		case "Duration": // type Duration int64
			f.typ = tInt64
			f.pkgReq = x.Name
			f.aliasType = tTimeDuration
			f.isDef = true
			f.isFixedLen = o.isLenFixed(f.typ)
			f.elmSize = o.isLen(f.typ)
			return true
		case "Time": // type Time struct
			f.typ = tTime
			f.pkgReq = x.Name
			f.aliasType = tTime
			f.isFixedLen = o.isLenFixed(f.typ)
			f.elmSize = o.isLen(f.typ)
			return true
		}
	}

	var imp string
	imp, ok = matchImport(fileImports, x.Name)
	if !ok {
		return
	}

	var err error
	f.typ, err = resolveImportedTypes(imp, d.Sel.Name)
	if err != nil {
		return false
	}

	f.pkgReq = imp
	f.aliasType = pkgSelName(x.Name, d.Sel.Name)
	f.isDef = true
	f.isFixedLen = o.isLenFixed(f.typ)
	f.elmSize = o.isLen(f.typ)
	return
}

func matchImport(fileImports []*dst.ImportSpec, pkg string) (importPath string, ok bool) {
	for _, fileImport := range fileImports {
		if fileImport.Path == nil || fileImport.Path.Value == "" {
			continue
		}

		importPath = strings.Trim(fileImport.Path.Value, "\"")
		if pkg == filepath.Base(importPath) {
			return importPath, true
		}
	}
	return "", false
}

func pkgSelName(pkg, selector string) string {
	if pkg != "" {
		return fmt.Sprintf("%s.%s", pkg, selector)
	}
	return selector
}

func packageName(f *dst.File) string {
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
	case *dst.BasicLit:
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

	case *dst.ValueSpec:
		if len(d.Values) == 1 {
			return calcArraySize(d.Values[0])
		}
	case *dst.Ident:
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

func nestedName(name *dst.Ident, parents ...[]*dst.Ident) (n string) {
	var list []string
	for _, parent := range append(parents, []*dst.Ident{name}) {
		if len(parent) != 1 {
			// TODO needs improving - expected only one parent.
			continue
		}
		list = append(list, parent[0].Name)
	}
	return strings.Join(list, ".")
}

func (s *structTyp) addExportedFields(names []*dst.Ident, f *field, parents [][]*dst.Ident) {
	f.structTyp = s

	for m := range names {
		f.name = nestedName(names[m], parents...)

		if f.typ == tBool || f.arrayType == tBool && f.isArray() {
			f.fieldList = &s.bool
			s.bool = append(s.bool, f)
			continue
		}
		switch f.typ {
		case tByte, tInt8:
			f.fieldList = &s.single
			s.single = append(s.single, f)
			continue
		case tStrings:
			f.fieldList = &s.stringSlice
			s.stringSlice = append(s.stringSlice, f)
			continue
		}
		// TODO add support for adding tiny enums using <= 7 bits

		if f.isFixedLen {
			f.fieldList = &s.fixedLen
			s.fixedLen = append(s.fixedLen, f)
		} else {
			f.fieldList = &s.variableLen
			s.variableLen = append(s.variableLen, f)
		}
	}
}

// isLen returns how many bytes each type requires.
func (o Option) isLen(typ string) uint {
	switch typ {
	case tBool, tByte, tInt8, tString:
		return 1
	case tInt16, tUint16:
		return 2
	case tInt32, tUint32, tFloat32:
		return 4
	case tInt64, tUint64, tFloat64, tTime, tComplex64:
		return 8
	case tComplex128:
		return 16
	case tInt, tUint:
		if o.Is32bit {
			return 4
		}
		return 8
	}
	return 0
}

func (o Option) isLenFixed(typ string) bool {
	switch typ {
	case tInt:
		return !o.VariableIntSize
	case tString:
		return false
	case tUint:
		return !o.VariableUintSize
	}
	return true
}

func bufWriteF(b *bytes.Buffer, format string, a ...any) {
	_, err := fmt.Fprintf(b, format, a...)
	if err != nil {
		lg.Println("err writing to buffer:", err)
	}
}

func bufWriteLine(b *bytes.Buffer, line string) {
	b.WriteByte('\t')
	b.WriteString(line)
	b.WriteByte('\n')
}

func bufWriteLineF(b *bytes.Buffer, format string, a ...any) {
	b.WriteByte('\t')
	bufWriteF(b, format, a...)
	b.WriteByte('\n')
}

// firstVarLenField return the first field that is variable length,
// either from s.stringSlice, otherwise s.variableLen.
// Returns nil when none are found.
func (s *structTyp) firstVarLenField() *field {
	if len(s.stringSlice) >= 1 {
		return s.stringSlice[0]
	}
	if len(s.variableLen) >= 1 {
		return s.variableLen[0]
	}
	return nil
}

func lenVariable(index uint) string {
	return lengthNameDefault + utl.UtoA(index)
}

// List of supported types.
const (
	tBool, tBools                            = "bool", "[]bool"
	tByte, tBytes                            = "byte", "[]byte"
	tComplex64, tComplex64s                  = "complex64", "[]complex64"
	tComplex128, tComplex128s                = "complex128", "[]complex128"
	tFloat32, tFloat32s                      = "float32", "[]float32"
	tFloat64, tFloat64s                      = "float64", "[]float64"
	tInt, tInt8, tInt16, tInt32, tInt64      = "int", "int8", "int16", "int32", "int64"
	tInts, tInt8s, tInt16s, tInt32s, tInt64s = "[]int", "[]int8", "[]int16", "[]int32", "[]int64"
	tUint, tUint16, tUint32, tUint64         = "uint", "uint16", "uint32", "uint64"
	tUints, tUint16s, tUint32s, tUint64s     = "[]uint", "[]uint16", "[]uint32", "[]uint64"
	tString, tStrings                        = "string", "[]string"
	tTime, tTimes                            = "time.Time", "[]time.Time"
	tTimeDuration, tTimeDurations            = "time.Duration", "[]time.Duration"
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
