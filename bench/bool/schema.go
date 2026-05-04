package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
)

type TwentyFour struct {
	One         bool
	Two         bool
	Three       bool
	Four        bool
	Five        bool
	Six         bool
	Seven       bool
	Eight       bool
	Nine        bool
	Ten         bool
	Eleven      bool
	Twelve      bool
	Thirteen    bool
	Fourteen    bool
	Fifteen     bool
	Sixteen     bool
	Seventeen   bool
	Eighteen    bool
	Nineteen    bool
	Twenty      bool
	TwentyOne   bool
	TwentyTwo   bool
	TwentyThree bool
	TwentyFour  bool
}

var (
	Obj24     = TwentyFour{One: true, Five: true, Nine: true, Ten: true, Twelve: true, Seventeen: true, Twenty: true, TwentyThree: true, TwentyFour: true}
	Bebop24   []byte
	Gob24     []byte
	Json24, _ = json.Marshal(Obj24)
	Jay24     = Obj24.MarshalJ()
)

func init() {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(Obj24)
	if err != nil {
		log.Fatalln(err)
	}
	Gob24 = buf.Bytes()

	buf.Reset()
	err = Obj24.EncodeBebop(&buf)
	if err != nil {
		log.Fatalln(err)
	}
	Bebop24 = buf.Bytes()
}
