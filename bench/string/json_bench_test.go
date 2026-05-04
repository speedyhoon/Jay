package main_test

import (
	"encoding/json"
	"testing"

	"github.com/speedyhoon/jay/bench/string"
)

var (
	output []byte
	err    error
)

func BenchmarkJSON_string_Marshal(b *testing.B) {
	for b.Loop() {
		output, err = json.Marshal(main.Obj24)
	}
}

func BenchmarkJSON_string_Unmarshal(b *testing.B) {
	var target main.TwentyFour
	for b.Loop() {
		err = json.Unmarshal(main.Json24, &target)
	}
}
