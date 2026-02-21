package string_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

// var str1 string
// var y = make([]byte, 19)
// var z = []byte("\022octopus camouflage")

func BenchmarkReadString67(b *testing.B) {
	for b.Loop() {
		str1, integer, bool1 = jay.ReadString(z)
	}
}

func BenchmarkReadStringPtr76(b *testing.B) {
	for b.Loop() {
		integer, bool1 = jay.ReadStringPtr(z, &str1)
	}
}
