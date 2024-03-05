package fuzz_test

import (
	"errors"
	"github.com/speedyhoon/jay/generate"
	"github.com/speedyhoon/rando/types"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const tempPath = "tmp_fuzz"

func init() {
	err := os.MkdirAll(tempPath, os.ModePerm)
	if err != nil {
		log.Panicln("Error creating directory:", err)
		return
	}
}

func TestFuzz(t *testing.T) {
	const perm = 0666

	pathPkg := filepath.Join(tempPath, "pkg.go")
	pathTest := filepath.Join(tempPath, "jay_test.go")
	pathJay := filepath.Join(tempPath, generate.DefaultOutputFileName)

	opt := generate.Option{FixedIntSize: true, FixedUintSize: true}

	pkg, tests, err := types.Package(tempPath)
	// Ensure both files are saved before processing. But if pathPkg fails to save, at least try to save pathTest too.
	err1 := os.WriteFile(pathPkg, pkg, perm)
	err2 := os.WriteFile(pathTest, tests, perm)
	require.NoError(t, errors.Join(err, err1, err2))

	require.NoError(t, opt.ProcessWrite(pkg, pathJay))

	cmd := exec.Command("go", "test")
	cmd.Dir = tempPath
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	err = cmd.Run()
	require.NoError(t, err)
}
