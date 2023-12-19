package main

import (
	"log"
	"time"
)

type Supported struct {
	Bool    bool
	Byte    byte
	Float32 float32
	Float64 float64
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Rune    rune
	String  string
	//Guid [16]byte
	Time time.Time
	Nano time.Time `j:"nano"`
	//Embed embed
	embed Embed
	Embed
	subStruct
	SubStruct subStruct
}

type Embed struct {
	Bool    bool
	Byte    byte
	Float32 float32
	Float64 float64
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Rune    rune
	//String  string
	//Guid [16]byte
	Time time.Time
	Nano time.Time `j:"nano"`
	//Embedded Supported
}

type subStruct struct {
	Bool    bool
	Byte    byte
	Float32 float32
	Float64 float64
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Rune    rune
	//String  string
	//Guid [16]byte
	Time time.Time
	Nano time.Time `j:"nano"`
	//Embedded Supported
}

func main() {
	log.SetFlags(log.Lshortfile)
}
