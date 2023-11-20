package jay

import (
	"go/ast"
	"go/token"
	"log"
	"reflect"
	"strings"
	"unicode"
)

const (
	StructTagName = "j"
	IgnoreFlag    = "-"
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
	const tagSymbol = '`'
	if len(str) >= 2 && str[0] == tagSymbol && str[len(str)-1] == tagSymbol {
		return str[1 : len(str)-1]
	}
	return str
}

func (s *Struct) ReceiverName() string {
	return string(unicode.ToLower([]rune(s.name)[0]))
}

func supportedType(typ string) bool {
	switch typ {
	case "bool", "byte" /*"float32","float64",*/, "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64":
		return true
	}
	return false
}

func (o Option) isSupportedType(t *ast.Field) (typeOf, aliasTypeName string, isVarLen, ok bool) {
	switch d := t.Type.(type) {
	case *ast.Ident:
		if supportedType(d.Name) {
			return d.Name, d.Name, o.isLenVariable(d.Name), true
		}

		// Type has an alias name.
		typeOf, isVarLen = o.determineType(d.Obj)
		return typeOf, d.Name, isVarLen, typeOf != ""
		//log.Printf("names: %s, type: %s, tag: %s\n", strings.Join(names, ","), d.Name, tag)
	case nil:
		// Ignore.
	default:
		log.Printf("type %T not expected here", d)
	}
	return "", "", false, false
}

func (o Option) determineType(t interface{}) (s string, isVarLen bool) {
	switch x := t.(type) {
	case *ast.Object:
		if x.Name == "" || x.Kind != ast.Typ {
			return "", false
		}
		return o.determineType(x.Decl)
	case *ast.TypeSpec:
		return o.determineType(x.Type)
	case *ast.StructType:
		if x.Fields != nil && len(x.Fields.List) != 0 {
			return "struct", o.hasVariableLen(x.Fields.List)
		}
	case nil:
		// Ignore.
	default:
		log.Printf("type %T not expected in determineType", x)
	}
	return "", false
}

func (o Option) hasVariableLen(fields []*ast.Field) bool {
	for _, f := range fields {
		if !hasExported(f.Names) {
			continue
		}

		_, _, isVarLen, ok := o.isSupportedType(f)
		if !ok {
			continue
		}

		if isVarLen {
			return true
		}
	}
	return false
}

func hasExported(idents []*ast.Ident) bool {
	for _, ident := range idents {
		if ident.IsExported() {
			return true
		}
	}
	return false
}

func (s *Struct) addExportedFields(names []*ast.Ident, tag, typeOf, aliasType string, isVarLen bool) {
	for m := range names {
		f := field{name: names[m].Name, tag: tag, typ: typeOf, aliasType: aliasType}
		f.LoadTagOptions()
		if f.typ == "bool" {
			s.bool = append(s.bool, f)
			continue
		}
		// TODO add support for adding tiny enums using <= 7 bits

		if isVarLen {
			s.variableLen = append(s.variableLen, f)
		} else {
			s.fixedLen = append(s.fixedLen, f)
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
	// TODO "float32","float64",
	case "bool", "byte", "int8", "uint8", "string", "int", "uint":
		return 1
	case "int16", "uint16":
		return 2
	case "int32", "rune", "uint32":
		return 4
	case "int64", "uint64":
		return 8
	}
	return 0
}

func (o Option) isLenVariable(typ string) bool {
	switch typ {
	// TODO case "[]byte", "[]bool", "[]int", "map[x]x",
	case "int":
		return !o.FixedIntSize
	case "string":
		return true
	case "uint":
		return !o.FixedUintSize
	}
	return false
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
