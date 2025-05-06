package boolDef

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/speedyhoon/rando"
	"log"
	"testing"
)

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

func Benchmark23_Marshal_Gob(b *testing.B) {
	var buf bytes.Buffer
	for b.Loop() {
		err = gob.NewEncoder(&buf).Encode(from23)
		output = buf.Bytes()
	}
}

func Benchmark23_Unmarshal_Gob(b *testing.B) {
	for b.Loop() {
		err = gob.NewDecoder(bytes.NewBuffer(bytes23Gob)).Decode(&to23)
	}
}

func Benchmark23_Marshal_Jay(b *testing.B) {
	for b.Loop() {
		output = from23.MarshalJ()
	}
}

func Benchmark23_Unmarshal_Jay(b *testing.B) {
	for b.Loop() {
		err = to23.UnmarshalJ(bytes23)
	}
}

func Benchmark23_Marshal_JSON(b *testing.B) {
	for b.Loop() {
		output, err = json.Marshal(from23)
	}
}

func Benchmark23_Unmarshal_JSON(b *testing.B) {
	for b.Loop() {
		err = json.Unmarshal(bytes23Json, &to23)
	}
}
