package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"mvdan.cc/gofumpt/format"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ProcessFile(filename string, source interface{}, opts ...Option) (src []byte, err error) {
	if filename == "" && source == nil {
		return nil, errors.New("no filename or source provided")
	}

	if filename != "" && !isGoFileName(filename) {
		return nil, fmt.Errorf("`%s` does not contain a Go file extension", filename)
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

	opt := LoadOptions(opts...)

	var list []Struct
	ast.Walk(visitor{structs: &list, option: opt}, f)

	src, err = generateFile(f.Name.Name, list, opt)
	if len(src) == 0 {
		return
	}

	// Nicely format the generated Go code.
	var bb []byte
	bb, err = format.Source(src, format.Options{
		LangVersion: strings.TrimPrefix(runtime.Version(), "go"),
		ExtraRules:  true,
	})
	if err != nil {
		return bb, err
	}
	return bb, err
}

// ProcessWrite processes a file and writes to outputFile.
func ProcessWrite(filename string, source interface{}, outputFile string, opts ...Option) (err error) {
	src, err := ProcessFile(filename, source, opts...)
	if err != nil || len(src) == 0 {
		return err
	}

	if outputFile == "" {
		outputFile = DefaultOutputFileName
	}

	err = os.WriteFile(filepath.Join(filepath.Dir(filename), outputFile), src, 0666)
	if err != nil {
		return err
	}

	return
}

func (s *Struct) Process(o Option, fields []*ast.Field) (hasExportedFields bool) {
	for i := 0; i < len(fields); i++ {
		t := fields[i]

		tag := getTag(t.Tag)
		if tag == IgnoreFlag {
			fields = Remove(fields, i)
			continue
		}

		typeOf, aliasType, isVarLen, ok := o.isSupportedType(t)
		if !ok {
			fields = Remove(fields, i)
			continue
		}

		names := getNames(t)
		if len(names) == 0 {
			fields = Remove(fields, i)
			continue
		}

		s.addExportedFields(names, tag, typeOf, aliasType, isVarLen)
	}

	return s.hasExportedFields()
}

func (s *Struct) hasExportedFields() bool {
	return len(s.fixedLen) >= 1 || len(s.variableLen) >= 1 || len(s.bool) >= 1
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

func isGoFileName(path string) bool {
	return strings.HasSuffix(strings.ToLower(path), goExt)
}
