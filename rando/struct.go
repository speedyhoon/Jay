package rando

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"mvdan.cc/gofumpt/format"
	"runtime"
	"strings"
)

// Package generates a Go main package with 1 to 6 randomly generated structs.
func Package(name string) (pkg, test []byte, _ error) {
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
	mx := rand.Intn(5) + 1
	for i := 0; i < mx; i++ {
		typeName := structNames.add(StringExported)
		structLines, testLines := Struct(typeName, 4000)
		b.Write(structLines)

		tests.WriteString(fmt.Sprintf(`
func TestFuzz_%d(t *testing.T) {
	var expected, actual %s
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, %[2]s{}, expected)
	require.Equal(t, %[2]s{}, actual)

	actual = %[2]s{
		%[3]s
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, %[2]s{}, expected)
	require.NotEqual(t, %[2]s{}, actual)
	require.Equal(t, expected, actual)
}
`, i, typeName, testLines))
	}

	pkg, err1 := goFormat(b.Bytes())
	test, err2 := goFormat(tests.Bytes())

	return pkg, test, errors.Join(err1, err2)
}

func goFormat(src []byte) (code []byte, err error) {
	code, err = format.Source(src, format.Options{
		LangVersion: strings.TrimPrefix(runtime.Version(), "go"),
		ExtraRules:  true,
	})
	if err != nil {
		return src, err
	}
	return
}

type uniqueNames map[string]struct{}

// add executes function f until a unique string is generated, not already in uniqueNames.
func (u uniqueNames) add(f func() string) string {
	for i := 0; i <= 999; i++ {
		tmp := f()
		_, ok := u[tmp]
		if !ok {
			u[tmp] = struct{}{}
			return tmp
		}
	}
	panic("function f does no generate enough randomness")
}

// Struct generates a random struct with a random number of fields.
func Struct(name string, fieldsQty uint) (fields, testLines []byte) {
	b := bytes.NewBufferString(fmt.Sprintf("type %s struct{\n", name))
	tl := bytes.NewBuffer(nil)

	fieldNames := make(uniqueNames)

	mx := rand.Intn(int(fieldsQty)) + 1
	for i := 0; i < mx; i++ {
		//b.WriteString(ExportedNames())
		n := fieldNames.add(FieldName)
		typ := Type()
		b.WriteString(n)
		b.WriteString("\t")
		b.WriteString(typ)
		b.WriteString("\n")

		// Test lines.
		tl.WriteString(n)
		tl.WriteString(":\trando.")
		tl.WriteString(strings.ToUpper(string(typ[0])))
		tl.WriteString(typ[1:])
		tl.WriteString("(),\n")
	}
	b.WriteString("}\n")
	return b.Bytes(), tl.Bytes()
}

func Type() string {
	supportedTypes := []string{
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
		//"struct",
		//"time.Time",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
	}
	return supportedTypes[rand.Intn(len(supportedTypes))]
}

func ExportedNames() string {
	fieldNames := make([]string, rand.Intn(5)+1)
	for j := range fieldNames {
		fieldNames[j] = StringExported()
	}
	return strings.Join(fieldNames, ", ")
}

func StringExported() string {
	s := UppercaseChar()

	mx := rand.Intn(100)
	for i := 0; i < mx; i++ {
		s += AlphaNumeric()
	}
	return s
}

func PackageName() string {
	s := LowercaseChar()

	mx := rand.Intn(100)
	for i := 0; i < mx; i++ {
		s += AlphaNumeric()
	}
	return s
}
func FieldName() string {
	s := AlphaChar()

	mx := rand.Intn(100)
	for i := 0; i < mx; i++ {
		s += AlphaNumeric()
	}
	return s
}

// AlphaNumeric returns an uppercase, lowercase letter or number digit.
func AlphaNumeric() string {
	switch rand.Intn(3) {
	case 1:
		return UppercaseChar()
	case 2:
		return LowercaseChar()
	default:
		return Number()
	}
}

// AlphaChar returns an uppercase or lowercase letter.
func AlphaChar() string {
	if rand.Intn(1) == 0 {
		return UppercaseChar()
	}
	return LowercaseChar()
}
func UppercaseChar() string {
	return string(rand.Intn('Z'-'A') + 'A')
}
func LowercaseChar() string {
	return string(rand.Intn('z'-'a') + 'a')
}
func Number() string {
	return string(rand.Intn('9'-'0') + '0')
}
