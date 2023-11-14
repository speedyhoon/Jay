package jay

import (
	"bytes"
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

func ProcessFile(filename string, source interface{}) (src []byte, err error) {
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

	var list []Struct
	ast.Walk(visitor{structs: &list}, f)

	return GenerateFile(f.Name.Name, Process(list))
}

func ProcessWrite(filename string, source interface{}) (err error) {
	// TODO support processing a whole directory
	src, err := ProcessFile(filename, source)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(filepath.Dir(filename), "jay.go"), src, 0666)
	if err != nil {
		return err
	}

	return
}

func Process(s []Struct) (src []byte) {
	buf := bytes.NewBuffer(nil)
	for i := 0; i < len(s); i++ {
		if !ast.IsExported(s[i].name) || len(s[i].fields) == 0 {
			s = Remove(s, i)
			continue
		}

		for j := 0; j < len(s[i].fields); j++ {
			t := s[i].fields[j]

			tag := getTag(t.Tag)
			if tag == "-" {
				// Tag has the ignore flag set.
				s[i].fields = Remove(s[i].fields, j)
				continue
			}

			typeOf, typeName, isVarLen, ok := isSupportedType(t)
			if typeOf == "struct" {
				log.Println()
			}
			if !ok {
				s[i].fields = Remove(s[i].fields, i)
				continue
			}

			name := t.Names
			if t.Names == nil {
				tt, ok := t.Type.(*ast.Ident)
				if !ok || tt == nil {
					s[i].fields = Remove(s[i].fields, i)
					continue
				}

				name = []*ast.Ident{tt}
			}

			s[i].addExportedFields(name, tag, typeOf, typeName, isVarLen)
		}
		s[i].GenerateFuncs(buf)
	}

	return buf.Bytes()
}

func Remove[T any](s []T, index int) []T {
	if index <= -1 {
		return s
	}

	l := len(s)
	if l <= 0 || index < 0 || index >= l {
		return s
	}

	switch index {
	case 0:
		return s[1:]
	case l - 1:
		return s[:index]
	default:
		return append(s[:index], s[index+1:]...)
	}
}
