package string

import (
	"github.com/speedyhoon/jay"
	"testing"
)

//const strSizeOf = 1

var str1 string
var integer int
var bool1 bool
var y = make([]byte, 19)
var z = []byte("\022octopus camouflage")

func BenchmarkReadString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str1, integer, bool1 = jay.ReadString(z)
	}
}

func BenchmarkReadStringV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str1, integer = ReadStringV2(z)
	}
}

func BenchmarkReadStringPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integer, bool1 = jay.ReadStringPtr(z, &str1)
	}
}

func BenchmarkReadStringPtrV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integer = ReadStringPtrV2(z, &str1)
	}
}

func ReadStringV2(b []byte) (h string, size int) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf || len(b) < size {
		return
	}

	return string(b[strSizeOf:size]), size
}

func ReadStringPtrV2(b []byte, h *string) (size int) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf || len(b) < size {
		return
	}

	*h = string(b[strSizeOf:size])
	return
}

var zebra Zebra
var err error

func BenchmarkReadStringPtrErrV1(b *testing.B) {
	//at := 0
	//var ok bool
	for i := 0; i < b.N; i++ {
		err = zebra.Unmarshal1(z)
		//c.Gearbox.Manufacturer, at, ok = jay.ReadString(b[at:])
		//if !ok {
		//	return jay.ErrUnexpectedEOB
		//}
	}
}
func BenchmarkReadStringPtrErrV2(b *testing.B) {
	//var err error
	for i := 0; i < b.N; i++ {
		//err = jay.ReadStringPtrErr(b[at:], &c.Gearbox.Manufacturer)
		err = zebra.Unmarshal2(z)
	}
}

type Zebra struct{ name string }

func (z *Zebra) Unmarshal1(b []byte) error {
	var ok bool
	z.name, _, ok = jay.ReadString(b)
	if !ok {
		return jay.ErrUnexpectedEOB
	}

	return nil
}
func (z *Zebra) Unmarshal2(b []byte) error {
	return jay.ReadStringPtrErr(b, &z.name)
}
