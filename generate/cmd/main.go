package main

import (
	"flag"
	"github.com/speedyhoon/jay/generate"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TODO fix error:
// go run ./cmd/main.go -o generate/tests/boolJay.go generate/tests/
//	main.go:73: open generate/tests/generate/tests/boolJay.go: no such file or directory
// but this works:
// go run ./cmd/main.go -o boolJay.go generate/tests/

// TODO fix error
// output file is not created in same directory as filepaths
// cd jay
// go run cmd/main.go -o jay.go bench/byte/make.go
// jay.go should be created in bench/byte/jay.go instead.

func main() {
	var opt generate.Option
	var outputFile string
	var types slice
	var verbose bool

	flag.BoolVar(&opt.Is32bit, "32", generate.IntSize == 32, "Force 32-bit output for ints & uints. Defaults to this systems 32-bit or 64-bit architecture.")
	flag.BoolVar(&opt.FixedIntSize, "fi", true, "Fixed int size.")
	flag.BoolVar(&opt.FixedUintSize, "fu", true, "Fixed uint size.")
	flag.StringVar(&outputFile, "o", generate.DefaultOutputFileName, "Output file.")
	flag.BoolVar(&verbose, "v", false, "Verbose output.")
	flag.BoolVar(&opt.SearchTests, "s", false, "Search Go test files for exported structs too.")
	flag.BoolVar(&opt.SkipTests, "t", false, "Don't generate Go test files.")
	flag.BoolVar(&opt.SkipMarshal, "m", false, "Don't generate MarshalJ() function.")
	flag.BoolVar(&opt.SkipUnmarshal, "u", false, "Don't generate UnmarshalJ() function.")
	flag.Var(&types, "y", "Exclusive list of comma separated types to generate marshalling and/or unmarshalling for. Default is to process all exported types. For example: Vet,animal.Cat,animal.Cow will process locally defined 'Vet' along with 'Cat' & 'Cow' in imported package \"animal\".")
	flag.Parse()
	paths := generate.RemoveDuplicates(flag.Args())

	if opt.SkipMarshal && opt.SkipUnmarshal {
		log.Println("Nothing to do. Both -m and -u flags are set.")
	}

	if verbose {
		opt.Verbose = log.New(os.Stdout, "", log.Lshortfile)
	} else {
		opt.Verbose = log.New(io.Discard, "", 0)
	}

	if len(types) >= 1 {
		opt.Verbose.Println("-y", types)
		opt.OnlyTypes = types
	}

	if len(paths) == 0 {
		paths = []string{"."}
	}

	var filePaths []string

	for _, path := range paths {
		path = filepath.Clean(path)
		if path == "" {
			path = "."
		}

		if isDir(path) {
			filePaths = append(filePaths, walkDir(path, opt)...)
		} else {
			filePaths = append(filePaths, path)
		}
	}

	err := opt.ProcessWrite(nil, outputFile, filePaths...)
	if err != nil {
		opt.Verbose.Println(err)
	}
}

func walkDir(path string, opt generate.Option) (filenames []string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && generate.IsGoFileName(info.Name()) {
			if !opt.SearchTests && generate.IsGoTestFileName(path) {
				opt.Verbose.Println("ignoring test file", path)
				return nil
			}

			filenames = append(filenames, path)
		}
		return nil
	})
	if err != nil {
		opt.Verbose.Println(err)
	}

	return
}

func isDir(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}

	fs, err := f.Stat()
	if err != nil {
		return false
	}

	return fs.IsDir()
}

type slice []string

func (l *slice) String() string {
	return strings.Join(*l, ", ")
}

func (l *slice) Set(s string) error {
	for _, item := range strings.Split(s, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			*l = append(*l, item)
		}
	}

	return nil
}
