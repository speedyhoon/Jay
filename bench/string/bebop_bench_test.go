package main_test

import (
	"bytes"
	"testing"

	"github.com/speedyhoon/jay/bench/string"
)

// go install github.com/200sc/bebop/main/bebopc-go
// bebopc-go -i schema.bop  -o bebop.go

func BenchmarkBebop_string_Marshal(b *testing.B) {
	buf := bytes.NewBuffer(nil)
	b.ResetTimer()

	for b.Loop() {
		err = main.Obj24.EncodeBebop(buf)
	}
}

func BenchmarkBebop_string_Unmarshal(b *testing.B) {
	var target main.TwentyFour
	buf := bytes.NewBuffer(main.Bebop24)
	b.ResetTimer()

	for b.Loop() {
		err = target.DecodeBebop(buf)
	}
}
