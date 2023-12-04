package main

import (
	"flag"
	"github.com/speedyhoon/jay/generate"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(log.Lshortfile)
	var opt generate.Option
	var outputFile string

	flag.BoolVar(&opt.Is32bit, "32", generate.IntSize == 32, "Force 32-bit output for ints & uints.")
	flag.BoolVar(&opt.FixedIntSize, "fi", true, "Fixed int size.")
	flag.BoolVar(&opt.FixedUintSize, "fu", true, "Fixed uint size.")
	flag.StringVar(&outputFile, "o", generate.DefaultOutputFileName, "Output file.")
	flag.Parse()
	paths := flag.Args()

	if len(paths) == 0 {
		paths = []string{"."}
	}

	for _, path := range paths {
		path = filepath.Clean(path)
		if path == "" {
			path = "."
		}

		if isDir(path) {
			walkDir(path, outputFile, opt)
		} else {
			process(path, outputFile, opt)
		}
	}
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

func walkDir(path, out string, opt generate.Option) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			process(path, out, opt)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func process(path, out string, opt generate.Option) {
	err := generate.ProcessWrite(path, nil, out, opt)
	if err != nil {
		log.Println(err)
	}
}
