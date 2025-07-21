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

	utl.DelDup(&filenames)
	for i := range filenames {
		if filenames[i] == o.outputFile {
			// Don't bother parsing the output file, as it will be overwritten anyway.
			continue
		}

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

	directories.walk(o)

	return o.makeFiles(directories)
}

func (d *dirList) walk(o *Option) {
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
			lg.Printf("makeFile err: %s,\ndirectory: %s,\npackage: %s\nsrc:\n%s\n", err, fl.pkg, dir, src)
		}

		if err == nil || o.IsDebug {
			// Write to disk when success or when the debug flag is set.
			output = append(output, Output{Dir: dir, Src: src})
		}
	}
	return
}

type (
	fileList struct {
		pkg     string
		structs []*structTyp
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
	if outputFile == "" {
		outputFile = DefaultOutputFileName
	}
	o.outputFile = outputFile

	output, err := o.ProcessFiles(source, filenames...)
	if err != nil {
		return err
	}

	if len(output) == 0 {
		return nil
	}

	dir, _ := filepath.Split(outputFile)
	var path string
	for i := range output {
		if dir == "" {
			path = filepath.Clean(filepath.Join(output[i].Dir, outputFile))
		} else {
			path = outputFile
		}
		err = os.WriteFile(path, output[i].Src, 0666)
		errors.Join(err)
	}

	return err
}

func (s *structTyp) process(fields []*ast.Field, dirList *dirList) (hasExportedFields bool) {
	for i := uint(0); i < utl.Len(fields); {
		t := fields[i]

		tag := getTag(t.Tag)
		if tag == IgnoreFlag {
			utl.Del(&fields, i)
			continue
		}

		if fs, ok := isStruct(t); ok {
			s.process(fs, dirList)
			i++
			continue
		}

		names := getNames(t)
		if len(names) == 0 {
			utl.Del(&fields, i)
			continue
		}

		fe := newField(tag)
		ok := s.option.isSupportedType(&fe, t.Type, dirList, s.dir)
		if !ok {
			utl.Del(&fields, i)
			continue
		}

		s.addExportedFields(names, &fe)
		// Only increment `i` if the field was added. If the field was removed, then `i` will still point to the next field.
		i++
	}

	s.setFieldByteIndexes()
	return s.hasExportedFields()
}

func structFields(t any, names []string) (fields []*ast.Field, ok bool) {
	switch d := t.(type) {
	case *ast.Ident:
		if d.Obj != nil {
			return structFields(d.Obj, names)
		}
	case *ast.Object:
		if d.Decl != nil {
			return structFields(d.Decl, names)
		}
	case *ast.TypeSpec:
		if d.Type != nil {
			return structFields(d.Type, names)
		}
	case *ast.StructType:
		if d.Fields != nil && len(d.Fields.List) >= 1 {
			prefixFieldList(d.Fields, names)
			return d.Fields.List, true
		}
	}
	return
}

func prefixFieldList(f *ast.FieldList, names []string) {
	qty := len(names)
	for i := range f.List {
		for j := range f.List[i].Names {
			if f.List[i].Names[j].IsExported() {

				// Duplicate the FieldList name for 2nd onwards names.
				for k := 1; k < qty; k++ {
					n := ast.Ident{
						Name: pkgSelName(names[k], f.List[i].Names[j].Name),
						Obj:  f.List[i].Names[j].Obj,
					}
					f.List[i].Names = append(f.List[i].Names, &n)
				}

				// Prefix the name of the first ident.
				f.List[i].Names[j].Name = pkgSelName(names[0], f.List[i].Names[j].Name)
			}
		}
	}
}

func isStruct(f *ast.Field) (fields []*ast.Field, ok bool) {
	names := getNames(f)
	if len(names) >= 1 {
		list := make([]string, len(names))
		for i, name := range names {
			list[i] = name.Name
		}
		return structFields(f.Type, list)
	}

	return
}

func (s *structTyp) setFirstNLast() {
	lists := s.processOrder()
	for i := range lists {
		if len(lists[i]) >= 1 {
			lists[i][0].isFirst = true
			goto setLast
		}
	}

setLast:
	// Reverse order loop.
	for i := len(lists) - 1; i >= 0; i-- {
		if n := len(lists[i]) - 1; n >= 0 {
			lists[i][n].isLast = true
			return
		}
	}
}

// processOrder returns is the order that fieldList's are processed in.
func (s *structTyp) processOrder() []fieldList {
	return []fieldList{s.bool, s.single, s.fixedLen, s.stringSlice, s.variableLen}
}

func (s *structTyp) setFieldByteIndexes() {
	var isFirstVarLen bool
	var byteIndex uint
	for _, f := range append(s.stringSlice, s.variableLen...) {
		if !f.isArray() {
			f.qtyIndex = []uint{byteIndex}
			byteIndex++
		}
	}

	for i, list := range s.processOrder() {
		var boolsQty uint
		for n, f := range list {
			si := byteIndex
			switch i {
			case 0: //bool
				si = byteIndex + boolsQty/8
				f.indexStart = &si
				boolsQty += f.totalSize()
				ei := byteIndex + (boolsQty-1)/8
				f.indexEnd = &ei
				if n == len(s.bool)-1 {
					byteIndex += (boolsQty + 7) >> 3
				}

			case 1: //single
				f.indexStart, f.indexEnd = &si, &si
				byteIndex++
			case 2: //fixedLen
				f.indexStart = &si
				byteIndex += f.totalSize()
				ei := byteIndex
				f.indexEnd = &ei
			case 3, 4: // stringSlice, variableLen
				if !isFirstVarLen {
					f.indexStart = &si
					isFirstVarLen = true
				}
			}
		}
	}

	s.qtyBytesRequired = byteIndex

	for _, f := range s.stringSlice {
		if f.isArray() {
			s.qtyBytesRequired += uint(f.arraySize)
		}
	}
}

func (f *field) totalSize() uint {
	if f.isArray() {
		return uint(f.arraySize) * f.elmSize
	}
	return f.elmSize
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
			utl.Del(&names, uint(i))
			continue
		}
		i++
	}

	return names
}
