package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

const (
	MaxUint8  = 1<<8 - 1  // 255
	MaxUint16 = 1<<16 - 1 // 65535
	MaxUint24 = 1<<24 - 1 // 16777215
	MaxUint32 = 1<<32 - 1 // 4294967295
	MaxUint40 = 1<<40 - 1 // 1099511627775
	MaxUint48 = 1<<48 - 1 // 281474976710655
	MaxUint56 = 1<<56 - 1 // 72057594037927935
	MaxUint64 = 1<<64 - 1 // 18446744073709551615

	MaxInt8  = 1<<7 - 1  // 127
	MinInt8  = -1 << 7   // -128
	MaxInt16 = 1<<15 - 1 // 32767
	MinInt16 = -1 << 15  // -32768
	MaxInt24 = 1<<23 - 1 // 8388607
	MinInt24 = -1 << 23  // -8388608
	MaxInt32 = 1<<31 - 1 // 2147483647
	MinInt32 = -1 << 31  // -2147483648
	MaxInt40 = 1<<39 - 1 // 549755813887
	MinInt40 = -1 << 39  // -549755813888
	MaxInt48 = 1<<47 - 1 // 140737488355327
	MinInt48 = -1 << 47  // -140737488355328
	MaxInt56 = 1<<55 - 1 // 36028797018963967
	MinInt56 = -1 << 55  // -36028797018963968
	MaxInt64 = 1<<63 - 1 // 9223372036854775807
	MinInt64 = -1 << 63  // -9223372036854775808
)

func TestWriteInt32(t *testing.T) {
	b := make([]byte, 5)
	jay.WriteInt32(b, MinInt32)
	assert.Equal(t, []byte{0, 0, 0, 128, 0}, b)
	jay.WriteInt32(b, MinInt24)
	assert.Equal(t, []byte{0, 0, 128, 255, 0}, b)
	jay.WriteInt32(b, MinInt16)
	assert.Equal(t, []byte{0, 128, 255, 255, 0}, b)
	jay.WriteInt32(b, MinInt8)
	assert.Equal(t, []byte{128, 255, 255, 255, 0}, b)
	jay.WriteInt32(b, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0}, b)
	jay.WriteInt32(b, MaxInt8)
	assert.Equal(t, []byte{127, 0, 0, 0, 0}, b)
	jay.WriteInt32(b, MaxUint8)
	assert.Equal(t, []byte{255, 0, 0, 0, 0}, b)
	jay.WriteInt32(b, MaxInt16)
	assert.Equal(t, []byte{255, 127, 0, 0, 0}, b)
	jay.WriteInt32(b, MaxUint16)
	assert.Equal(t, []byte{255, 255, 0, 0, 0}, b)
	jay.WriteInt32(b, MaxInt24)
	assert.Equal(t, []byte{255, 255, 127, 0, 0}, b)
	jay.WriteInt32(b, MaxUint24)
	assert.Equal(t, []byte{255, 255, 255, 0, 0}, b)
	jay.WriteInt32(b, MaxInt32)
	assert.Equal(t, []byte{255, 255, 255, 127, 0}, b)
}

func TestWriteInt16(t *testing.T) {
	b := make([]byte, 3)
	jay.WriteInt16(b, MinInt16)
	assert.Equal(t, []byte{0, 128, 0}, b)
	jay.WriteInt16(b, MinInt8)
	assert.Equal(t, []byte{128, 255, 0}, b)
	jay.WriteInt16(b, 0)
	assert.Equal(t, []byte{0, 0, 0}, b)
	jay.WriteInt16(b, MaxInt8)
	assert.Equal(t, []byte{127, 0, 0}, b)
	jay.WriteInt16(b, MaxUint8)
	assert.Equal(t, []byte{255, 0, 0}, b)
	jay.WriteInt16(b, MaxInt16)
	assert.Equal(t, []byte{255, 127, 0}, b)
}

func TestWriteInt64(t *testing.T) {
	b := make([]byte, 8)
	jay.WriteInt64(b, MinInt64)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 128}, b)
	jay.WriteInt64(b, MinInt56)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 128, 255}, b)
	jay.WriteInt64(b, MinInt48)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 128, 255, 255}, b)
	jay.WriteInt64(b, MinInt40)
	assert.Equal(t, []byte{0, 0, 0, 0, 128, 255, 255, 255}, b)
	jay.WriteInt64(b, MinInt32)
	assert.Equal(t, []byte{0, 0, 0, 128, 255, 255, 255, 255}, b)
	jay.WriteInt64(b, MinInt24)
	assert.Equal(t, []byte{0, 0, 128, 255, 255, 255, 255, 255}, b)
	jay.WriteInt64(b, MinInt16)
	assert.Equal(t, []byte{0, 128, 255, 255, 255, 255, 255, 255}, b)
	jay.WriteInt64(b, MinInt8)
	assert.Equal(t, []byte{128, 255, 255, 255, 255, 255, 255, 255}, b)
	jay.WriteInt64(b, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxInt8)
	assert.Equal(t, []byte{127, 0, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxUint8)
	assert.Equal(t, []byte{255, 0, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxInt16)
	assert.Equal(t, []byte{255, 127, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxUint16)
	assert.Equal(t, []byte{255, 255, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxInt24)
	assert.Equal(t, []byte{255, 255, 127, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxUint24)
	assert.Equal(t, []byte{255, 255, 255, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxInt32)
	assert.Equal(t, []byte{255, 255, 255, 127, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxUint32)
	assert.Equal(t, []byte{255, 255, 255, 255, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxInt40)
	assert.Equal(t, []byte{255, 255, 255, 255, 127, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxUint40)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 0, 0, 0}, b)
	jay.WriteInt64(b, MaxInt48)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 127, 0, 0}, b)
	jay.WriteInt64(b, MaxUint48)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 0, 0}, b)
	jay.WriteInt64(b, MaxInt56)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 127, 0}, b)
	jay.WriteInt64(b, MaxUint56)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 0}, b)
	jay.WriteInt64(b, MaxInt64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 127}, b)
}

func TestWriteIntNew(t *testing.T) {
	b := make([]byte, 8)
	for i := MinInt8 - 1; i <= MaxInt8+1; i++ {
		j := lenInt(i)
		jay.WriteIntVariable(b, i, j)
		if i == 128 {
			println("")
		}
		r, l := jay.ReadInt(b)
		assert.Equal(t, i, r)
		if i == 0 {
			assert.Equal(t, j, l)
		} else {
			assert.Equal(t, j, l-1)
		}
	}
}

func TestInt(t *testing.T) {
	b := make([]byte, 9)

	maxList := []int{
		0,
		MinInt64, MinInt56, MinInt48, MinInt40,
		MinInt32, MinInt24, MinInt16, MinInt8,
		MaxInt8, MaxInt16, MaxInt24, MaxInt32,
		MaxInt40, MaxInt48, MaxInt56, MaxInt64,
	}
	for n := range maxList {
		j := lenInt(maxList[n])
		if maxList[n] == MaxInt40 {
			println("")
		}
		jay.WriteIntVariable(b, maxList[n], j)
		r, l := jay.ReadInt(b)

		if maxList[n] == 0 {
			assert.Equal(t, j, l)
		} else {
			assert.Equal(t, j, l-1)
		}
		assert.Equalf(t, maxList[n], r, "index: %d", n)
	}
}

func TestWriteInt(t *testing.T) {
	b := make([]byte, 9)
	jay.WriteIntVariable(b, MinInt64, lenInt(MinInt64))
	assert.Equal(t, []byte{8, 0, 0, 0, 0, 0, 0, 0, 128}, b)
	b = make([]byte, 8)
	jay.WriteIntVariable(b, MinInt56, lenInt(MinInt56))
	assert.Equal(t, []byte{7, 0, 0, 0, 0, 0, 0, 128}, b)
	b = make([]byte, 7)
	jay.WriteIntVariable(b, MinInt48, lenInt(MinInt48))
	assert.Equal(t, []byte{6, 0, 0, 0, 0, 0, 128}, b)
	b = make([]byte, 6)
	jay.WriteIntVariable(b, MinInt40, lenInt(MinInt40))
	assert.Equal(t, []byte{5, 0, 0, 0, 0, 128}, b)
	b = make([]byte, 5)
	jay.WriteIntVariable(b, MinInt32, lenInt(MinInt32))
	assert.Equal(t, []byte{4, 0, 0, 0, 128}, b)
	b = make([]byte, 4)
	jay.WriteIntVariable(b, MinInt24, lenInt(MinInt24))
	assert.Equal(t, []byte{3, 0, 0, 128}, b)
	b = make([]byte, 3)
	jay.WriteIntVariable(b, MinInt16, lenInt(MinInt16))
	assert.Equal(t, []byte{2, 0, 128}, b)
	b = make([]byte, 2)
	jay.WriteIntVariable(b, MinInt8, lenInt(MinInt8))
	assert.Equal(t, []byte{1, 128}, b)
	b = make([]byte, 1)
	jay.WriteIntVariable(b, 0, lenInt(0))
	assert.Equal(t, []byte{0}, b)
	b = make([]byte, 2)
	jay.WriteIntVariable(b, MaxInt8, lenInt(MaxInt8))
	assert.Equal(t, []byte{1, 127}, b)
	b = make([]byte, 3)
	jay.WriteIntVariable(b, MaxUint8, lenInt(MaxUint8))
	assert.Equal(t, []byte{2, 255, 0}, b)
	jay.WriteIntVariable(b, MaxInt16, lenInt(MaxInt16))
	assert.Equal(t, []byte{2, 255, 127}, b)
	b = make([]byte, 4)
	jay.WriteIntVariable(b, MaxUint16, lenInt(MaxUint16))
	assert.Equal(t, []byte{3, 255, 255, 0}, b)
	jay.WriteIntVariable(b, MaxInt24, lenInt(MaxInt24))
	assert.Equal(t, []byte{3, 255, 255, 127}, b)
	b = make([]byte, 5)
	jay.WriteIntVariable(b, MaxUint24, lenInt(MaxUint24))
	assert.Equal(t, []byte{4, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, MaxInt32, lenInt(MaxInt32))
	assert.Equal(t, []byte{4, 255, 255, 255, 127}, b)
	b = make([]byte, 6)
	jay.WriteIntVariable(b, MaxUint32, lenInt(MaxUint32))
	assert.Equal(t, []byte{5, 255, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, MaxInt40, lenInt(MaxInt40))
	assert.Equal(t, []byte{5, 255, 255, 255, 255, 127}, b)
	b = make([]byte, 7)
	jay.WriteIntVariable(b, MaxUint40, lenInt(MaxUint40))
	assert.Equal(t, []byte{6, 255, 255, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, MaxInt48, lenInt(MaxInt48))
	assert.Equal(t, []byte{6, 255, 255, 255, 255, 255, 127}, b)
	b = make([]byte, 8)
	jay.WriteIntVariable(b, MaxUint48, lenInt(MaxUint48))
	assert.Equal(t, []byte{7, 255, 255, 255, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, MaxInt56, lenInt(MaxInt56))
	assert.Equal(t, []byte{7, 255, 255, 255, 255, 255, 255, 127}, b)
	b = make([]byte, 9)
	jay.WriteIntVariable(b, MaxUint56, lenInt(MaxUint56))
	assert.Equal(t, []byte{8, 255, 255, 255, 255, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, MaxInt64, lenInt(MaxInt64))
	assert.Equal(t, []byte{8, 255, 255, 255, 255, 255, 255, 255, 127}, b)
}

func TestReadInt24(t *testing.T) {
	b := make([]byte, 3)
	// Increment each iteration by between 75 and 100.
	inc := rand.Intn(25) + 75
	for i := MinInt24; i <= MaxInt24; i += inc {
		jay.WriteInt24(b, i)
		require.Equal(t, i, jay.ReadInt24(b), b)
	}
}

/*
// ReadInt24 ...
func ReadInt24(b []byte) int {
	//if b[0]&_128 == 0 {

	value := int(b[0]) | int(b[1])<<8 | int(b[2])<<16
	return value | -(value & 0x00800000) // Apply a sign extension if the highest bit is set
	//}

	//return | neg24Mask int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16
	///}
	//return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16
}
*/

func TestReadInt40(t *testing.T) {
	b := make([]byte, 5)
	inc := MaxUint32 - rand.Intn(150)
	for i := MinInt40; i <= MaxInt40; i += inc {
		jay.WriteInt40(b, i)
		require.Equal(t, i, jay.ReadInt40(b), b)
	}
}

func TestReadInt48(t *testing.T) {
	b := make([]byte, 6)
	inc := MaxUint32 - rand.Intn(150)
	for i := MinInt48; i <= MaxInt48; i += inc {
		jay.WriteInt48(b, i)
		require.Equal(t, i, jay.ReadInt48(b), b)
	}
}

func TestReadInt56(t *testing.T) {
	b := make([]byte, 7)
	inc := MaxUint48 - rand.Intn(150)
	for i := MinInt56; i <= MaxInt56; i += inc {
		jay.WriteInt56(b, i)
		require.Equal(t, i, jay.ReadInt56(b), b)
	}
}

func TestReadInt64(t *testing.T) {
	assert.Equal(t, int64(0), jay.ReadInt64([]byte{0, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, int64(247), jay.ReadInt64([]byte{247, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, int64(MaxInt64), jay.ReadInt64([]byte{255, 255, 255, 255, 255, 255, 255, 127}))
}

func TestReadInt32(t *testing.T) {
	assert.Equal(t, int32(0), jay.ReadInt32([]byte{0, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, int32(247), jay.ReadInt32([]byte{247, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, int32(-1), jay.ReadInt32([]byte{255, 255, 255, 255, 255, 255, 255, 127}))
	assert.Equal(t, int32(MaxInt32), jay.ReadInt32([]byte{255, 255, 255, 127, 255, 255, 255, 127}))
}

func lenInt(i int) int {
	if i == 0 {
		return 0
	}

	if i > MaxInt56 || i < MinInt56 {
		return 8
	}

	if i > MaxInt48 || i < MinInt48 {
		return 7
	}

	if i > MaxInt40 || i < MinInt40 {
		return 6
	}

	if i > MaxInt32 || i < MinInt32 {
		return 5
	}

	if i > MaxInt24 || i < MinInt24 {
		return 4
	}

	if i > MaxInt16 || i < MinInt16 {
		return 3
	}

	if i > MaxInt8 || i < MinInt8 {
		return 2
	}

	return 1
}
