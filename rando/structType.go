package rando

import (
	"bytes"
	"errors"
	"fmt"
)

// PackageTypes generates a Go main package a random struct with the given field types.
// It returns the package and test code required to test MarshalJ() & UnmarshalJ().
func PackageTypes(name string, types []string) (pkg, test []byte, err error) {
	b := bytes.NewBufferString(fmt.Sprintf("package %s\n\n", name))
	structNames := make(uniqueNames)

	tests := bytes.NewBuffer(b.Bytes())
	tests.WriteString(`

import (
	"github.com/speedyhoon/jay/rando"
	"github.com/stretchr/testify/require"
	"testing"
)

`)
	structs := bytes.NewBuffer(nil)
	typeName := structNames.add(StringExported)
	structLines, testLines := StructType(typeName, types)
	structs.Write(structLines)

	tests.WriteString(fmt.Sprintf(`
func TestFuzz_1(t *testing.T) {
	var expected, actual %[1]s
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, %[1]s{}, expected)
	require.Equal(t, %[1]s{}, actual)

	actual = %[1]s{
		%[2]s
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	//require.NotEqual(t, %[1]s{}, expected)
	//require.NotEqual(t, %[1]s{}, actual)
	require.Equal(t, expected, actual)
}
`, typeName, testLines))

	if bytes.Contains(structs.Bytes(), []byte("time.Time")) {
		b.WriteString("import \"time\"\n\n")
	}
	b.Write(structs.Bytes())

	pkg, err1 := goFormat(b.Bytes())
	test, err2 := goFormat(tests.Bytes())

	return pkg, test, errors.Join(err1, err2)
}

// StructType generates a struct with random field names with the given types.
func StructType(name string, types []string) (fields, testLines []byte) {
	b := bytes.NewBufferString(fmt.Sprintf("type %s struct{\n", name))
	tl := bytes.NewBuffer(nil)

	fieldNames := make(uniqueNames)

	for _, typ := range types {
		//b.WriteString(ExportedNames())
		n := fieldNames.add(FieldName)
		b.WriteString(n)
		b.WriteString("\t")
		b.WriteString(typ)
		b.WriteString("\n")

		// Test lines.
		tl.WriteString(n)
		tl.WriteString(":\t")
		tl.WriteString(TypeRandomFunc(typ))
		tl.WriteString(",\n")
	}
	b.WriteString("}\n")
	return b.Bytes(), tl.Bytes()
}
