package byte

import "testing"

var int1 int
var y = make([]byte, 19)

func BenchmarkWriteBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		write2(y, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7}, 18)
	}
}

func BenchmarkWriteBytes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7], y[8], y[9], y[10], y[11], y[12], y[13], y[14], y[15], y[16], y[17], y[18] = 19, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18
	}
}
func BenchmarkWriteBytes2B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7], y[8], y[9], y[10], y[11], y[12], y[13], y[14], y[15], y[16], y[17], y[18] = 19, array[0], array[1], array[2], array[3], array[4], array[5], array[6], array[7], array[8], array[9], array[10], array[11], array[12], array[13], array[14], array[15], array[16], array[17]
	}
}
func BenchmarkWriteBytes2C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y[0] = byte(18) // Set how long the string is.
		copy(y[1:], array[:])
	}
}
func BenchmarkWriteBytes3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y[0] = 19
		y[1] = 1
		y[2] = 2
		y[3] = 3
		y[4] = 4
		y[5] = 5
		y[6] = 6
		y[7] = 7
		y[8] = 8
		y[9] = 9
		y[10] = 10
		y[11] = 11
		y[12] = 12
		y[13] = 13
		y[14] = 14
		y[15] = 15
		y[16] = 16
		y[17] = 17
		y[18] = 18
	}
}

var array = [18]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7}
var empty []byte

func BenchmarkWriteBytes4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		write3(y, array[:], 18)
	}
}
func BenchmarkWriteBytes5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		write3F(y, array[:], 18)
	}
}
func BenchmarkWriteBytes4B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		write3(y, empty, 0)
	}
}
func BenchmarkWriteBytes5B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		write3F(y, empty, 0)
	}
}

func BenchmarkWriteBytes4At(b *testing.B) {
	for i := 0; i < b.N; i++ {
		int1 = writeAt(y, array[:], 18, 0)
	}
}
func BenchmarkWriteBytes5At(b *testing.B) {
	for i := 0; i < b.N; i++ {
		int1 = writeAtF(y, array[:], 18, 0)
	}
}
func BenchmarkWriteBytes4AtB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		int1 = writeAt(y, empty, 0, 0)
	}
}
func BenchmarkWriteBytes5AtB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		int1 = writeAtF(y, empty, 0, 0)
	}
}
