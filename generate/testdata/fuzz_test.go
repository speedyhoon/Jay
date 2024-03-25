package testdata_test

import (
	"errors"
	"github.com/speedyhoon/jay/generate"
	"github.com/speedyhoon/rando/types"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestFuzz(t *testing.T) {
	for _, typ := range []string{
		"bool",
		"byte",
		"float32",
		"float64",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"rune",
		"string",
		//"struct{}",
		"time.Time",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"[]byte",
		"[]uint8",
	} {
		tempPath := strings.ReplaceAll(strings.Trim(typ, "[].{}"), ".", "")
		if strings.HasPrefix(typ, "[]") {
			tempPath += "s"
		}
		GenerateFuzzTest(t, tempPath, typ)
	}
}

func GenerateFuzzTest(t *testing.T, fuzzDir string, typ string) {
	err := os.MkdirAll(fuzzDir, os.ModePerm)
	if err != nil {
		log.Panicln("Error creating directory:", err)
		return
	}

	const perm = 0666

	pathPkg := filepath.Join(fuzzDir, "pkg.go")
	pathTest := filepath.Join(fuzzDir, "jay_test.go")
	pathJay := filepath.Join(fuzzDir, generate.DefaultOutputFileName)

	opt := generate.Option{FixedIntSize: true, FixedUintSize: true, PointerMarshalFunc: true}

	pkg, tests, err := types.PackageSequence("main", typ, "testdata/fuzz_test")
	// Ensure both files are saved before processing. But if pathPkg fails to save, at least try to save pathTest too.
	err1 := os.WriteFile(pathPkg, pkg, perm)
	err2 := os.WriteFile(pathTest, tests, perm)
	require.NoError(t, errors.Join(err, err1, err2))

	require.NoError(t, opt.ProcessWrite(pkg, pathJay))

	cmd := exec.Command("go", "test")
	cmd.Dir = fuzzDir
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	err = cmd.Run()
	require.NoError(t, err)
}
