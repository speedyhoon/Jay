package rando

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"mvdan.cc/gofumpt/format"
	"runtime"
	"strings"
)

// PackageMain generates a Go main package with 1 to 6 randomly generated structs.
func PackageMain() []byte {
	b := bytes.NewBufferString("package main\n\n")
	mainFunc := bytes.NewBufferString("func main() {\n")
	structNames := make(uniqueNames)

	define := ":"
	mx := rand.Intn(5) + 1
	for i := 0; i < mx; i++ {
		typeName := structNames.add(StringExported)
		b.Write(Struct(typeName, 4000))

		varName := i + 'a'
		mainFunc.WriteString(fmt.Sprintf(`var %[1]c %[2]s
	err %[3]s= %[1]c.UnmarshalJ(%[1]c.MarshalJ())
	if err != nil {
		panic(err)
	}
`, varName, typeName, define))
		define = ""
	}

	mainFunc.WriteString("}")

	code := append(b.Bytes(), mainFunc.Bytes()...)

	// Pretty format generated package code.
	src, err := format.Source(code, format.Options{
		LangVersion: strings.TrimPrefix(runtime.Version(), "go"),
		ExtraRules:  true,
	})
	if err != nil {
		log.Panicf("rando.Package generated errornous code:\n%s\n", code)
	}

	return src
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
func Struct(name string, fieldsQty uint) []byte {
	b := bytes.NewBufferString(fmt.Sprintf("type %s struct{\n", name))

	fieldNames := make(uniqueNames)

	mx := rand.Intn(int(fieldsQty)) + 1
	for i := 0; i < mx; i++ {
		//b.WriteString(ExportedNames())
		b.WriteString(fieldNames.add(FieldName))
		b.WriteString("\t")
		b.WriteString(TypeOf())
		b.WriteString("\n")
	}
	b.WriteString("}\n")
	return b.Bytes()
}

func TypeOf() string {
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
