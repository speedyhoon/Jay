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

const (
	goExt      = ".go"
	testSuffix = "_test" + goExt
)

var verbose = log.New(io.Discard, "", log.Lshortfile)

func main() {
	log.SetFlags(log.Lshortfile)
	var opt generate.Option
	var outputFile string

	flag.BoolVar(&opt.Is32bit, "32", generate.IntSize == 32, "Force 32-bit output for ints & uints.")
	flag.BoolVar(&opt.FixedIntSize, "fi", true, "Fixed int size.")
	flag.BoolVar(&opt.FixedUintSize, "fu", true, "Fixed uint size.")
	flag.StringVar(&outputFile, "o", generate.DefaultOutputFileName, "Output file.")
	flag.BoolVar(&opt.Verbose, "v", false, "Verbose output.")
	flag.BoolVar(&opt.IncludeTests, "t", false, "Include Go test files.")
	flag.Parse()
	paths := flag.Args()

	if opt.Verbose {
		verbose.SetOutput(os.Stdout)
	}

	if len(paths) == 0 {
		paths = []string{"."}
	}

	// Batch files based on their directory location to prevent structs from being referenced in the wrong package (directory).
	dirBatch := make(map[string][]string)

	for _, path := range paths {
		path = filepath.Clean(path)
		if path == "" {
			path = "."
		}

		if isDir(path) {
			process(opt, filepath.Join(path, outputFile), walkDir(path, opt)...)
		} else {
			batchAppend(&dirBatch, path)
		}
	}

	for _, list := range dirBatch {
		// Process all files in the same directory so the package has the correct name and structures.
		process(opt, outputFile, list...)
	}
}

func batchAppend(filePaths *map[string][]string, file string) {
	dir := filepath.Dir(file)
	list, _ := (*filePaths)[dir]
	(*filePaths)[dir] = append(list, file)
}

func process(opt generate.Option, outputFile string, filePaths ...string) {
	err := opt.ProcessWrite(nil, outputFile, filePaths...)
	if err != nil {
		log.Println(err)
	}
}

func walkDir(path string, opt generate.Option) (filenames []string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && strings.HasSuffix(info.Name(), goExt) {
			if !opt.IncludeTests && strings.HasSuffix(path, testSuffix) {
				verbose.Println("ignoring test file", path)
				return nil
			}

			filenames = append(filenames, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
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
