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
	var f *ast.File
	var err error
	directories := make(dirList)

	if source != nil {
		f, err = ParseFile("", source)
		if err != nil {
			lg.Println("source error:", err)
			errors.Join(errs, err)
		} else {
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
			err = errors.Join(errs, err)
			// If a file is unable to be parsed, continue parsing the other files.
			continue
		}
		directories.add(filepath.Dir(filenames[i]), f)
	}

	directories.walk(*o)

	return o.makeFiles(directories)
}

func (d *dirList) walk(o Option) {
	for dir, fl := range *d {
		for _, file := range fl.files {
			ast.Walk(visitor{structs: &fl.structs, option: o, dir: dir, dirList: d}, file)
		}
		(*d)[dir] = fl
	}
}

func (o Option) makeFiles(directories dirList) (output []Output, errs error) {
	var src []byte
	var err error

	// Traverse the directories again because some imports weren't populated in the correct order to run makeFile() immediately after ast.Walk().
	for dir, fl := range directories {
		if len(fl.structs) == 0 {
			lg.Println("no exported structs in directory", dir)
			continue
		}

		src, err = o.makeFile(fl.pkg, fl.structs)
		if err != nil {
			errors.Join(errs, err)
			lg.Printf("makeFile err: %s, in dir: %s file:\n%s\n", err, dir, src)
		} else {
			output = append(output, Output{Dir: dir, Src: src})
		}
	}
	return
}

type (
	fileList struct {
		pkg     string
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
	if dir == "" {
		dir = "."
	}
	list, _ := (*d)[dir]
	if list.pkg == "" {
		list.pkg = packageName(file)
	}
	list.files = append(list.files, file)
	(*d)[dir] = list
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
		err = os.WriteFile(filepath.Clean(filepath.Join(output[i].Dir, outputFile)), output[i].Src, 0666)
		errors.Join(err)
	}

	return err
}

func (s *structTyp) process(fields []*ast.Field, o Option, dirList *dirList) (hasExportedFields bool) {
	for i := 0; i < len(fields); {
		t := fields[i]

		tag := getTag(t.Tag)
		if tag == IgnoreFlag {
			fields = Remove(fields, i)
			continue
		}

		fe, ok := o.isSupportedType(t.Type, dirList, s.dir)
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
		// Only increment `i` if the field was added. If the field was removed, `i` will still point to the next field.
		i++
	}

	s.setFirstNLast()
	return s.hasExportedFields()
}

func (s *structTyp) setFirstNLast() {
	// lists is the order that each field list is processed.
	lists := []fieldList{s.bool, s.boolArray, s.single, s.fixedLen, s.variableLen}
	for i := range lists {
		for n := range lists[i] {
			lists[i][n].isFirst = true
			goto setLast
		}
	}

setLast:
	// Reverse order loop.
	for i := len(lists) - 1; i >= 0; i-- {
		for n := len(lists[i]) - 1; n >= 0; n-- {
			lists[i][n].isLast = true
			return
		}
	}
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
