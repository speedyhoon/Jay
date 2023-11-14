package jay_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

func BenchmarkReadBool1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.ReadBool1(byte(i))
	}
}
func BenchmarkReadBool2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.ReadBool2(byte(i))
	}
}
func BenchmarkReadBool3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.ReadBool3(byte(i))
	}
}
func BenchmarkReadBool4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.ReadBool4(byte(i))
	}
}
func BenchmarkReadBool5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.ReadBool5(byte(i))
	}
}
func BenchmarkReadBool6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.ReadBool6(byte(i))
	}
}
func BenchmarkReadBool7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.ReadBool7(byte(i))
	}
}
func BenchmarkReadBool8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.ReadBool8(byte(i))
	}
}

func BenchmarkBool1True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool1(true)
	}
}
func BenchmarkBool1False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool1(false)
	}
}
func BenchmarkBool2True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool2(true, true)
	}
}
func BenchmarkBool2False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool2(false, false)
	}
}
func BenchmarkBool3True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool3(true, true, true)
	}
}
func BenchmarkBool3False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool3(false, false, false)
	}
}
func BenchmarkBool4True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool4(true, true, true, true)
	}
}
func BenchmarkBool4False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool4(false, false, false, false)
	}
}
func BenchmarkBool5True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool5(true, true, true, true, true)
	}
}
func BenchmarkBool5False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool5(false, false, false, false, false)
	}
}
func BenchmarkBool6True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool6(true, true, true, true, true, true)
	}
}
func BenchmarkBool6False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool6(false, false, false, false, false, false)
	}
}
func BenchmarkBool7True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool7(true, true, true, true, true, true, true)
	}
}
func BenchmarkBool7False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool7(false, false, false, false, false, false, false)
	}
}
func BenchmarkBool8True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool8(true, true, true, true, true, true, true, true)
	}
}
func BenchmarkBool8False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.Bool8(false, false, false, false, false, false, false, false)
	}
}
