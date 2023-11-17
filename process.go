package jay

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ProcessFile(filename string, source interface{}, opts ...Option) (src []byte, err error) {
	if filename == "" && source == nil {
		return nil, errors.New("no filename or source provided")
	}

	if filename != "" {
		if !strings.HasSuffix(filename, ".go") /*|| strings.HasPrefix(filename, ".")*/ {
			return nil, fmt.Errorf("`%s` does not contain a Go file extension", filename)
		}
	}

	//f, err := decorator.NewDecorator(token.NewFileSet()).ParseFile(filename, source, parser.ParseComments)
	//if err != nil {
	//	return
	//}

	//fSet := token.NewFileSet()
	f, err := parser.ParseFile(token.NewFileSet(), filename, source, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return
	}
	if f == nil {
		return nil, io.ErrUnexpectedEOF
	}

	//var opt Options
	//if len(options)>=1{
	//	opt=options[0]
	//	opt.Load()
	//}

	var list []Struct
	ast.Walk(visitor{structs: &list}, f)

	return GenerateFile(f.Name.Name, list, LoadOptions(opts...))
}

func ProcessWrite(filename string, source interface{}, opts ...Option) (err error) {
	// TODO support processing a whole directory
	src, err := ProcessFile(filename, source, opts...)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(filepath.Dir(filename), "jay.go"), src, 0666)
	if err != nil {
		return err
	}

	return
}

func (s *Struct) Process(fields []*ast.Field) (hasExportedFields bool) {
	for i := 0; i < len(fields); i++ {
		t := fields[i]

		tag := getTag(t.Tag)
		if tag == IgnoreFlag {
			fields = Remove(fields, i)
			continue
		}

		typeOf, typeName, isVarLen, ok := isSupportedType(t)
		if !ok {
			fields = Remove(fields, i)
			continue
		}

		names := getNames(t)
		if len(names) == 0 {
			fields = Remove(fields, i)
			continue
		}

		s.addExportedFields(names, tag, typeOf, typeName, isVarLen)
	}

	return len(s.fixedLen) >= 1 || len(s.variableLen) >= 1
}

func Remove[T any](t []T, index int) []T {
	if index < 0 {
		return t
	}

	l := len(t)
	if l == 0 || index >= l {
		return t
	}

	switch index {
	case 0:
		return t[1:]
	case l - 1:
		return t[:index]
	default:
		return append(t[:index], t[index+1:]...)
	}
}

func getNames(f *ast.Field) (names []*ast.Ident) {
	if len(f.Names) == 0 {
		idt, ok := f.Type.(*ast.Ident)
		if !ok || idt == nil {
			if !ok {
				log.Printf("unexpected type %T\n", f.Type)
			}
			return nil
		}

		return onlyExportedNames(idt)
	}

	return onlyExportedNames(f.Names...)
}

func onlyExportedNames(names ...*ast.Ident) []*ast.Ident {
	for i := 0; i < len(names); {
		if !ast.IsExported(names[i].Name) {
			names = Remove(names, i)
			continue
		}
		i++
	}

	return names
}
