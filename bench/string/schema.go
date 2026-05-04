package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
)

type TwentyFour struct {
	One         string
	Two         string
	Three       string
	Four        string
	Five        string
	Six         string
	Seven       string
	Eight       string
	Nine        string
	Ten         string
	Eleven      string
	Twelve      string
	Thirteen    string
	Fourteen    string
	Fifteen     string
	Sixteen     string
	Seventeen   string
	Eighteen    string
	Nineteen    string
	Twenty      string
	TwentyOne   string
	TwentyTwo   string
	TwentyThree string
	TwentyFour  string
}

var (
	Obj24     = TwentyFour{One: "Buffalo", Five: "Cheetah", Nine: "Elephant", Ten: "Giraffe", Twelve: "Hippopotamus", Seventeen: "Leopard", Twenty: "Lion", TwentyThree: "Rhinoceros", TwentyFour: "Zebras"}
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
