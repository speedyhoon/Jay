package buffer

import (
	"bytes"
	"io"
	"log"
	"testing"
)

var buf = bytes.NewBuffer(nil)

func BenchmarkBuffer(b *testing.B) {
	log.SetOutput(io.Discard)
	var i uint
	for b.Loop() {
		WriteUint(buf, i)
		i++
	}
	log.Println(buf.String())
}

//var buf2 = bytes.NewBuffer(nil)

func BenchmarkBuffer2(b *testing.B) {
	log.SetOutput(io.Discard)
	var buf1 = bytes.NewBuffer(nil)
	var i uint
	for b.Loop() {
		WriteUintBytes(buf1, i)
		i++
	}
	log.Println(buf1.String())
}

func BenchmarkBuffer23(b *testing.B) {
	log.SetOutput(io.Discard)
	var buf1 = bytes.NewBuffer(nil)
	var i uint
	for b.Loop() {
		WriteUintBytes2(buf1, i)
		i++
	}
	log.Println(buf1.String())
}
func BenchmarkBuffer3(b *testing.B) {
	log.SetOutput(io.Discard)
	var buf1 = bytes.NewBuffer(nil)
	var i uint
	for b.Loop() {
		WriteUintBytes3(buf1, i)
		i++
	}
	log.Println(buf1.String())
}

//func BenchmarkBytes(b *testing.B) {
//	g := make([]byte, 1)
//	var i uint
//	for b.Loop( {
//		PutUint32(g, i)
//		i++
//	}
//}

/*func BenchmarkBytes2(b *testing.B) {
	g := make([]byte, 1000000000*8)
	//var i uint
	for b.Loop() {
		PutUint32(g, 5)
		i++
	}
}*/
