package testdata_test

import (
	"errors"
	"github.com/speedyhoon/jay/generate"
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/assert"
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
	} {
		tempPath := strings.ReplaceAll(strings.Trim(typ, "[].{}"), ".", "")
		if strings.HasPrefix(typ, "[]") {
			tempPath += "s"
		}
		Genny(t, tempPath, typ)
	}
}

func Genny(t *testing.T, tempPath string, typ string) {
	err := os.MkdirAll(tempPath, os.ModePerm)
	if err != nil {
		log.Panicln("Error creating directory:", err)
		return
	}

	const perm = 0666

	pathPkg := filepath.Join(tempPath, "pkg.go")
	pathTest := filepath.Join(tempPath, "jay_test.go")
	pathJay := filepath.Join(tempPath, generate.DefaultOutputFileName)

	opt := generate.Option{FixedIntSize: true, FixedUintSize: true}

	pkg, tests, err := rando.PackageSequence("main", typ /*eList*/)
	// Ensure both files are saved before processing. But if pathPkg fails to save, at least try to save pathTest too.
	err1 := os.WriteFile(pathPkg, pkg, perm)
	err2 := os.WriteFile(pathTest, tests, perm)
	require.NoError(t, errors.Join(err, err1, err2))

	src, err := opt.ProcessFiles(pkg)
	assert.NoErrorf(t, err, "src: %s", src)
	require.NoError(t, os.WriteFile(pathJay, src, perm))

	cmd := exec.Command("go", "test")
	cmd.Dir = tempPath
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	err = cmd.Run()
	require.NoError(t, err)
}
