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

const testSuffix = "_test.go"

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

	var filePaths []string

	for _, path := range paths {
		path = filepath.Clean(path)
		if path == "" {
			path = "."
		}

		if isDir(path) {
			filePaths = append(filePaths, walkDir(path, outputFile, opt)...)
		} else {
			filePaths = append(filePaths, path)
		}
	}

	err := opt.ProcessWrite(nil, outputFile, filePaths...)
	if err != nil {
		log.Println(err)
	}
}

func walkDir(path, out string, opt generate.Option) (filenames []string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
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
