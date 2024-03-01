package generate

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Verbose = log.New(io.Discard, "", log.Lshortfile)

// ProcessFiles ...
func (o *Option) ProcessFiles(source interface{}, filenames ...string) (src []byte, err error) {
	if source == nil && len(filenames) == 0 {
		return nil, ErrNoSource
	}

	*o = LoadOptions(*o)
	var files []*ast.File
	var f *ast.File

	if source != nil {
		f, err = ParseFile("", source)
		if err != nil {
			log.Println("source error:", err)
		} else {
			files = append(files, f)
		}
	}

	for i := range filenames {
		if !isGoFileName(filenames[i]) {
			log.Printf("`%s` does not contain a Go file extension", filenames[i])
			continue
		}

		var fi os.FileInfo
		fi, err = os.Stat(filenames[i])
		if err != nil {
			log.Printf("unable to retrieve file info for %s: %s", filenames[i], err)
			continue
		}
		if fi.Size() == 0 {
			log.Printf("ignoring empty Go file %s", filenames[i])
			continue
		}

		f, err = ParseFile(filenames[i], nil)
		if err != nil {
			return
		}
		files = append(files, f)
	}

	var list []structTyp
	if len(files) == 0 {
		return
	}

	for i := range files {
		ast.Walk(visitor{structs: &list, option: *o}, files[i])
	}

	src, err = makeFile(files[0].Name.Name, list, *o)
	if err != nil {
		log.Println("err", err.Error())
	}
	return
}

func ParseFile(filename string, src interface{}) (f *ast.File, err error) {
	f, err = parser.ParseFile(token.NewFileSet(), filename, src, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return
	}
	if f == nil {
		return nil, io.ErrUnexpectedEOF
	}
	return
}

// ProcessWrite processes a file and writes to outputFile.
func (o *Option) ProcessWrite(source interface{}, outputFile string, filenames ...string) (err error) {
	src, err := o.ProcessFiles(source, filenames...)
	if err != nil || len(src) == 0 {
		return err
	}

	if outputFile == "" {
		outputFile = DefaultOutputFileName

		if len(filenames) != 0 {
			outputFile = filepath.Join(filepath.Dir(filenames[0]), outputFile)
		}
	}

	err = os.WriteFile(outputFile, src, 0666)
	return
}

func (s *structTyp) process(fields []*ast.Field, o Option) (hasExportedFields bool) {
	for i := 0; i < len(fields); {
		t := fields[i]

		tag := getTag(t.Tag)
		if tag == IgnoreFlag {
			fields = Remove(fields, i)
			continue
		}

		fe, ok := o.isSupportedType(t)
		if !ok {
			fields = Remove(fields, i)
			continue
		}

		names := getNames(t)
		if len(names) == 0 {
			fields = Remove(fields, i)
			continue
		}

		fe.tag = tag
		fe.LoadTagOptions()

		s.addExportedFields(names, fe)
		i++
	}

	return s.hasExportedFields()
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
