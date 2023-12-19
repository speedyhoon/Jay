package main

import "testing"

var car = Car{
	ID:      42,
	Row:     42389,
	Name:    "Hello",
	Auto:    true,
	RedLine: 65535,
	Gearbox: gearbox{
		Gears:        4,
		Reverse:      1,
		Sequential:   true,
		Automatic:    false,
		Model:        "H23w",
		Manufacturer: "Zebra",
	},
}
var src []byte
var ll int

func BenchmarkMarshalJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src = car.MarshalJ()
	}
}

//	func BenchmarkMarshalJBuf(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			src = car.MarshalJ2()
//		}
//	}
/*func BenchmarkMarshalJX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src = car.MarshalJX()
	}
}
func BenchmarkMarshalJY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src = car.MarshalJY()
	}
}

func BenchmarkMarshalK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src = car.MarshalK()
	}
}
func BenchmarkMarshalL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src = car.MarshalL()
	}
}*/

func BenchmarkLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ll = len(src)
	}
}

func BenchmarkFastest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fastest(i)
	}
}
func Fastest(i int) {
	// nothing to do here
	_ = i + 1
}
