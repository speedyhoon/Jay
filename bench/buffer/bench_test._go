package main

import (
	"bytes"
	"io"
	"log"
	"testing"
)

var buf = bytes.NewBuffer(nil)

func BenchmarkBuffer(b *testing.B) {
	log.SetOutput(io.Discard)
	var i, h uint = 0, uint(b.N)
	for ; i < h; i++ {
		WriteUint(buf, i)
	}
	log.Println(buf.String())
}

//var buf2 = bytes.NewBuffer(nil)

func BenchmarkBuffer2(b *testing.B) {
	log.SetOutput(io.Discard)
	var buf1 = bytes.NewBuffer(nil)
	var i, h uint = 0, uint(b.N)
	for ; i < h; i++ {
		WriteUintBytes(buf1, i)
	}
	log.Println(buf1.String())
}

func BenchmarkBuffer23(b *testing.B) {
	log.SetOutput(io.Discard)
	var buf1 = bytes.NewBuffer(nil)
	var i, h uint = 0, uint(b.N)
	for ; i < h; i++ {
		WriteUintBytes2(buf1, i)
	}
	log.Println(buf1.String())
}
func BenchmarkBuffer3(b *testing.B) {
	log.SetOutput(io.Discard)
	var buf1 = bytes.NewBuffer(nil)
	var i, h uint = 0, uint(b.N)
	for ; i < h; i++ {
		WriteUintBytes3(buf1, i)
	}
	log.Println(buf1.String())
}

//func BenchmarkBytes(b *testing.B) {
//	g := make([]byte, 1)
//	var i, h uint = 0, uint(b.N)
//	for ; i < h; i++ {
//		PutUint32(g, i)
//	}
//}

/*func BenchmarkBytes2(b *testing.B) {
	g := make([]byte, 1000000000*8)
	//var i, h uint = 0, uint(b.N)
	for i := 0; i < b.N; i++ {

		PutUint32(g, 5)
	}
}*/
