package fuzz_test

import (
	"github.com/speedyhoon/jay/generate"
	"github.com/speedyhoon/jay/rando"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const tempPath = "tmp"

func init() {
	err := os.MkdirAll(tempPath, os.ModePerm)
	if err != nil {
		log.Panicln("Error creating directory:", err)
		return
	}
}

func TestFuzz(t *testing.T) {
	//for i := 0; i < 10000; i++ {
	pkg := rando.PackageMain()
	opt := generate.Option{FixedIntSize: true, FixedUintSize: true}
	src, err := opt.ProcessFiles(pkg)
	require.NoErrorf(t, err, "src: %s", src)

	pathPkg, pathJay := filepath.Join(tempPath, "pkg.go"), filepath.Join(tempPath, "jay.go")

	// Create a temporary file.
	err = os.WriteFile(pathPkg, pkg, 0666)
	require.NoError(t, err)
	err = os.WriteFile(pathJay, src, 0666)
	require.NoError(t, err)

	cmd := exec.Command("go", "run", pathPkg, pathJay)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	err = cmd.Run()
	require.NoError(t, err)
}
