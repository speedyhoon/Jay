package main_test

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/speedyhoon/jay/bench/string"
)

func BenchmarkGob_string_Marshal(b *testing.B) {
	for b.Loop() {
		var buf bytes.Buffer
		err = gob.NewEncoder(&buf).Encode(main.Obj24)
	}
}

func BenchmarkGob_string_Unmarshal(b *testing.B) {
	for b.Loop() {
		var target main.TwentyFour
		err = gob.NewDecoder(bytes.NewReader(main.Gob24)).Decode(&target)
	}
}
