package boolDef

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/speedyhoon/rando"
	"log"
)
import "testing"

var to23 TwentyThree
var from23 = TwentyThree{
	One:         toggle(rando.Bool()),
	Four:        toggle(rando.Bool()),
	Five:        toggle(rando.Bool()),
	Six:         toggle(rando.Bool()),
	Seven:       toggle(rando.Bool()),
	Eight:       toggle(rando.Bool()),
	Nine:        toggle(rando.Bool()),
	Ten:         toggle(rando.Bool()),
	Eleven:      toggle(rando.Bool()),
	Twelve:      toggle(rando.Bool()),
	Thirteen:    toggle(rando.Bool()),
	Fourteen:    toggle(rando.Bool()),
	Fifteen:     toggle(rando.Bool()),
	Sixteen:     toggle(rando.Bool()),
	Seventeen:   toggle(rando.Bool()),
	Eighteen:    toggle(rando.Bool()),
	Nineteen:    toggle(rando.Bool()),
	Twenty:      toggle(rando.Bool()),
	TwentyOne:   toggle(rando.Bool()),
	TwentyTwo:   toggle(rando.Bool()),
	TwentyThree: toggle(rando.Bool()),
}

var bytes23 = from23.MarshalJ()
var bytes23Json, err = json.Marshal(from23)
var bytes23Gob []byte
var output []byte

func init() {
	fmt.Println("Benchmarking Unmarshal23: ", bytes23)

	var buf bytes.Buffer
	err = gob.NewEncoder(&buf).Encode(from23)
	if err != nil {
		log.Panic("gob encode error:", err)
	}
	bytes23Gob = buf.Bytes()
}

func Benchmark_Marshal23_Gob(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		err = gob.NewEncoder(&buf).Encode(from23)
		output = buf.Bytes()
	}
}

func Benchmark_Unmarshal23_Gob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = gob.NewDecoder(bytes.NewBuffer(bytes23Gob)).Decode(&to23)
	}
}

func Benchmark_Marshal23_J(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output = from23.MarshalJ()
	}
}

func Benchmark_Unmarshal23_J(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = to23.UnmarshalJ(bytes23)
	}
}

func Benchmark_Marshal23_JSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output, err = json.Marshal(&from23)
	}
}

func Benchmark_Unmarshal23_JSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(bytes23Json, &to23)
	}
}
