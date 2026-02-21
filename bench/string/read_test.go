package string_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

var integer int

func BenchmarkReadString(b *testing.B) {
	for b.Loop() {
		str1, integer, bool1 = jay.ReadString(z)
	}
}

func BenchmarkReadStringPtr(b *testing.B) {
	for b.Loop() {
		integer, bool1 = jay.ReadStringPtr(z, &str1)
	}
}
