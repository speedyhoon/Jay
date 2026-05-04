package main_test

import (
	"testing"

	"github.com/speedyhoon/jay/bench/bool"
)

func BenchmarkJay_bool_Marshal(b *testing.B) {
	for b.Loop() {
		output = main.Obj24.MarshalJ()
	}
}

func BenchmarkJay_bool_Unmarshal(b *testing.B) {
	var target main.TwentyFour
	for b.Loop() {
		err = target.UnmarshalJ(main.Jay24)
	}
}
