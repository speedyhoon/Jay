package jay_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

var integer int

func BenchmarkReadInt24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integer = jay.ReadInt24([]byte{240, 48, 77})
	}
}

func BenchmarkReadInt24v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integer = v1ReadInt24([]byte{240, 48, 77})
	}
}

func BenchmarkReadInt24v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integer = v2ReadInt24([]byte{240, 48, 77})
	}
}

func v1ReadInt24(b []byte) int {
	value := int(b[0]) | int(b[1])<<8 | int(b[2])<<16
	return value | -(value & 0x00800000) // Apply a sign extension if the highest bit is set
}

func v2ReadInt24(b []byte) int {
	return int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | -((int(b[0]) | int(b[1])<<8 | int(b[2])<<16) & 0x00800000) // Apply a sign extension if the highest bit is set
}
