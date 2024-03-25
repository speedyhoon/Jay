package generate

import (
	"errors"
	"github.com/speedyhoon/utl"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
)

var lg = log.New(io.Discard, "", log.Lshortfile)

// ProcessFiles ...
func (o *Option) ProcessFiles(source interface{}, filenames ...string) (output []Output, errs error) {
	if source == nil && len(filenames) == 0 {
		return nil, ErrNoSource
	}

	*o = LoadOptions(*o)
	var files []*ast.File
	var f *ast.File
	var err error
	directories := make(dirList)

	if source != nil {
		f, err = ParseFile("", source)
		if err != nil {
			lg.Println("source error:", err)
			errors.Join(errs, err)
		} else {
			files = append(files, f)
			directories.add(packageName(f), f)
		}
	}

	filenames = RemoveDuplicates(filenames)
	for i := range filenames {
		if !utl.IsGoFileName(filenames[i]) {
			lg.Printf("`%s` does not contain a Go file extension", filenames[i])
			continue
		}

		var fi os.FileInfo
		fi, err = os.Stat(filenames[i])
		if err != nil {
			errors.Join(errs, err)
			lg.Printf("unable to retrieve file info for %s: %s", filenames[i], err)
			continue
		}
		if fi.Size() == 0 {
			lg.Printf("ignoring empty Go file %s", filenames[i])
			continue
		}

		f, err = ParseFile(filenames[i], nil)
		if err != nil {
			errors.Join(errs, err)
			continue
		}
		files = append(files, f)
		directories.add(filepath.Dir(filenames[i]), f)
	}

	directories.Walk(*o)

	return o.makeFiles(directories)
}

func (d *dirList) Walk(o Option) {
	files := d.allFiles()
	for dir, fl := range *d {
		for _, file := range fl.files {
			ast.Walk(visitor{structs: &fl.structs, option: o, dir: dir, files: files}, file)
		}
		(*d)[dir] = fl
	}
}

func (o Option) makeFiles(directories dirList) (output []Output, errs error) {
	var src []byte
	var err error
	structList := directories.allStructs()

	// Traverse the directories again because some imports weren't populated in the correct order to run makeFile() immediately after ast.Walk().
	for dir, fl := range directories {
		if len(fl.structs) == 0 {
			lg.Println("no exported structs in directory", dir)
			continue
		}

		src, err = o.makeFile(filepath.Base(dir), structList)
		if err != nil {
			errors.Join(errs, err)
			lg.Println("makeFile:", err)
		} else {
			output = append(output, Output{Dir: dir, Src: src})
		}
	}
	return
}

type (
	fileList struct {
		structs []structTyp
		files   []*ast.File
	}
	dirList map[string]fileList
	Output  struct {
		Src []byte
		Dir string
	}
)

func (d *dirList) add(dir string, file *ast.File) {
	if dir == "." {
		if file != nil && file.Name != nil && file.Name.Name != "" {
			dir = file.Name.Name
		} else {
			dir = "main"
		}
	}
	list, _ := (*d)[dir]
	list.files = append(list.files, file)
	(*d)[dir] = list
}

func (d dirList) allStructs() (st []structTyp) {
	for _, dirs := range d {
		st = append(st, dirs.structs...)
	}
	return
}

func (d dirList) allFiles() (files []*ast.File) {
	for _, dirs := range d {
		files = append(files, dirs.files...)
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
	output, err := o.ProcessFiles(source, filenames...)
	if err != nil {
		return err
	}

	if len(output) == 0 {
		return nil
	}

	if outputFile == "" {
		outputFile = DefaultOutputFileName
	}

	for i := range output {
		err = os.WriteFile(outputFile, output[i].Src, 0666)
		errors.Join(err)
	}

	return err
}

func (s *structTyp) process(fields []*ast.Field, o Option, files []*ast.File) (hasExportedFields bool) {
	for i := 0; i < len(fields); {
		t := fields[i]

		tag := getTag(t.Tag)
		if tag == IgnoreFlag {
			fields = Remove(fields, i)
			continue
		}

		fe, ok := o.isSupportedType(t, files)
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
				lg.Printf("unexpected type %T\n", f.Type)
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
