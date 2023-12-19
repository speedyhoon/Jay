package string_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

var str1 string
var y = make([]byte, 19)
var z = []byte("\022octopus camouflage")
var bool1 bool

func BenchmarkWriteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.WriteString(y, "octopus camouflage", 18)
	}
}

func BenchmarkWriteStringEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.WriteString(y, "", 0)
	}
}

func BenchmarkWriteStringV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteStringV2(y, "octopus camouflage", 18)
	}
}

func BenchmarkWriteStringV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteStringV3(y, "octopus camouflage", 18)
	}
}

func BenchmarkWriteStringV2empty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteStringV2(y, "", 0)
	}
}

func WriteStringV2(y []byte, s string, length int) {
	y[0] = byte(length) // Set how long the string is.
	copy(y[1:length+1], s)
}

func WriteStringV3(y []byte, s string, length int) {
	y[0] = byte(length) // Set how long the string is.
	copy(y[1:], s)
}

func WriteStringV4(y []byte, s string, length int) {
	y[0] = byte(length) // Set how long the string is.
	if length != 0 {
		copy(y[strSizeOf:], s)
	}
}

const strSizeOf = 1

func WriteStringAtV2(y []byte, s string, length, at int) int {
	y[0] = byte(length) // Set how long the string is.
	if length != 0 {
		copy(y[strSizeOf:], s)
	}
	return at + length + strSizeOf
}

/*// Deprecated
func WriteStringOld(b []byte, index int, s *string) int {
	l := len(*s)
	b[index] = byte(l) // Set how long the string is
	index++
	copy(b[index:], *s)
	return index + l
}*/

// Deprecated
// WriteStringDeprecated expects b to be the start of a byte slice.
// TODO check if this will break when writing to byte slices that are too short!
// TODO check which function is faster - WriteString or WriteString2
func WriteStringDeprecated(b []byte, s string) (l int) {
	l = len(s)
	b[0] = byte(l) // Set how long the string is
	//index++
	if l != 0 {
		copy(b[1:l], s)
	}
	return l + 1
}
