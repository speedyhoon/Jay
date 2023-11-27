package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

const (
	MaxInt64 = 1<<63 - 1 // 9223372036854775807
	MinInt64 = -1 << 63  // -9223372036854775808
)

func TestWriteInt32(t *testing.T) {
	b := make([]byte, 5)
	jay.WriteInt32(b, jay.MinInt32)
	assert.Equal(t, []byte{0, 0, 0, 128, 0}, b)
	jay.WriteInt32(b, jay.MinInt24)
	assert.Equal(t, []byte{0, 0, 128, 255, 0}, b)
	jay.WriteInt32(b, jay.MinInt16)
	assert.Equal(t, []byte{0, 128, 255, 255, 0}, b)
	jay.WriteInt32(b, jay.MinInt8)
	assert.Equal(t, []byte{128, 255, 255, 255, 0}, b)
	jay.WriteInt32(b, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0}, b)
	jay.WriteInt32(b, jay.MaxInt8)
	assert.Equal(t, []byte{127, 0, 0, 0, 0}, b)
	jay.WriteInt32(b, jay.MaxUint8)
	assert.Equal(t, []byte{255, 0, 0, 0, 0}, b)
	jay.WriteInt32(b, jay.MaxInt16)
	assert.Equal(t, []byte{255, 127, 0, 0, 0}, b)
	jay.WriteInt32(b, jay.MaxUint16)
	assert.Equal(t, []byte{255, 255, 0, 0, 0}, b)
	jay.WriteInt32(b, jay.MaxInt24)
	assert.Equal(t, []byte{255, 255, 127, 0, 0}, b)
	jay.WriteInt32(b, jay.MaxUint24)
	assert.Equal(t, []byte{255, 255, 255, 0, 0}, b)
	jay.WriteInt32(b, jay.MaxInt32)
	assert.Equal(t, []byte{255, 255, 255, 127, 0}, b)
}

func TestWriteInt16(t *testing.T) {
	b := make([]byte, 3)
	jay.WriteInt16(b, jay.MinInt16)
	assert.Equal(t, []byte{0, 128, 0}, b)
	jay.WriteInt16(b, jay.MinInt8)
	assert.Equal(t, []byte{128, 255, 0}, b)
	jay.WriteInt16(b, 0)
	assert.Equal(t, []byte{0, 0, 0}, b)
	jay.WriteInt16(b, jay.MaxInt8)
	assert.Equal(t, []byte{127, 0, 0}, b)
	jay.WriteInt16(b, jay.MaxUint8)
	assert.Equal(t, []byte{255, 0, 0}, b)
	jay.WriteInt16(b, jay.MaxInt16)
	assert.Equal(t, []byte{255, 127, 0}, b)
}

func TestWriteInt64(t *testing.T) {
	b := make([]byte, 8)
	jay.WriteInt64(b, MinInt64)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 128}, b)
	jay.WriteInt64(b, jay.MinInt56)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 128, 255}, b)
	jay.WriteInt64(b, jay.MinInt48)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 128, 255, 255}, b)
	jay.WriteInt64(b, jay.MinInt40)
	assert.Equal(t, []byte{0, 0, 0, 0, 128, 255, 255, 255}, b)
	jay.WriteInt64(b, jay.MinInt32)
	assert.Equal(t, []byte{0, 0, 0, 128, 255, 255, 255, 255}, b)
	jay.WriteInt64(b, jay.MinInt24)
	assert.Equal(t, []byte{0, 0, 128, 255, 255, 255, 255, 255}, b)
	jay.WriteInt64(b, jay.MinInt16)
	assert.Equal(t, []byte{0, 128, 255, 255, 255, 255, 255, 255}, b)
	jay.WriteInt64(b, jay.MinInt8)
	assert.Equal(t, []byte{128, 255, 255, 255, 255, 255, 255, 255}, b)
	jay.WriteInt64(b, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxInt8)
	assert.Equal(t, []byte{127, 0, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxUint8)
	assert.Equal(t, []byte{255, 0, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxInt16)
	assert.Equal(t, []byte{255, 127, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxUint16)
	assert.Equal(t, []byte{255, 255, 0, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxInt24)
	assert.Equal(t, []byte{255, 255, 127, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxUint24)
	assert.Equal(t, []byte{255, 255, 255, 0, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxInt32)
	assert.Equal(t, []byte{255, 255, 255, 127, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxUint32)
	assert.Equal(t, []byte{255, 255, 255, 255, 0, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxInt40)
	assert.Equal(t, []byte{255, 255, 255, 255, 127, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxUint40)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 0, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxInt48)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 127, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxUint48)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 0, 0}, b)
	jay.WriteInt64(b, jay.MaxInt56)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 127, 0}, b)
	jay.WriteInt64(b, jay.MaxUint56)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 0}, b)
	jay.WriteInt64(b, MaxInt64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 127}, b)
}

func TestWriteIntNew(t *testing.T) {
	b := make([]byte, 8)
	for i := jay.MinInt8 - 1; i <= jay.MaxInt8+1; i++ {
		j := jay.LenInt(i)
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
		MinInt64, jay.MinInt56, jay.MinInt48, jay.MinInt40,
		jay.MinInt32, jay.MinInt24, jay.MinInt16, jay.MinInt8,
		jay.MaxInt8, jay.MaxInt16, jay.MaxInt24, jay.MaxInt32,
		jay.MaxInt40, jay.MaxInt48, jay.MaxInt56, MaxInt64,
	}
	for n := range maxList {
		j := jay.LenInt(maxList[n])
		if maxList[n] == jay.MaxInt40 {
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
	jay.WriteIntVariable(b, MinInt64, jay.LenInt(MinInt64))
	assert.Equal(t, []byte{8, 0, 0, 0, 0, 0, 0, 0, 128}, b)
	b = make([]byte, 8)
	jay.WriteIntVariable(b, jay.MinInt56, jay.LenInt(jay.MinInt56))
	assert.Equal(t, []byte{7, 0, 0, 0, 0, 0, 0, 128}, b)
	b = make([]byte, 7)
	jay.WriteIntVariable(b, jay.MinInt48, jay.LenInt(jay.MinInt48))
	assert.Equal(t, []byte{6, 0, 0, 0, 0, 0, 128}, b)
	b = make([]byte, 6)
	jay.WriteIntVariable(b, jay.MinInt40, jay.LenInt(jay.MinInt40))
	assert.Equal(t, []byte{5, 0, 0, 0, 0, 128}, b)
	b = make([]byte, 5)
	jay.WriteIntVariable(b, jay.MinInt32, jay.LenInt(jay.MinInt32))
	assert.Equal(t, []byte{4, 0, 0, 0, 128}, b)
	b = make([]byte, 4)
	jay.WriteIntVariable(b, jay.MinInt24, jay.LenInt(jay.MinInt24))
	assert.Equal(t, []byte{3, 0, 0, 128}, b)
	b = make([]byte, 3)
	jay.WriteIntVariable(b, jay.MinInt16, jay.LenInt(jay.MinInt16))
	assert.Equal(t, []byte{2, 0, 128}, b)
	b = make([]byte, 2)
	jay.WriteIntVariable(b, jay.MinInt8, jay.LenInt(jay.MinInt8))
	assert.Equal(t, []byte{1, 128}, b)
	b = make([]byte, 1)
	jay.WriteIntVariable(b, 0, jay.LenInt(0))
	assert.Equal(t, []byte{0}, b)
	b = make([]byte, 2)
	jay.WriteIntVariable(b, jay.MaxInt8, jay.LenInt(jay.MaxInt8))
	assert.Equal(t, []byte{1, 127}, b)
	b = make([]byte, 3)
	jay.WriteIntVariable(b, jay.MaxUint8, jay.LenInt(jay.MaxUint8))
	assert.Equal(t, []byte{2, 255, 0}, b)
	jay.WriteIntVariable(b, jay.MaxInt16, jay.LenInt(jay.MaxInt16))
	assert.Equal(t, []byte{2, 255, 127}, b)
	b = make([]byte, 4)
	jay.WriteIntVariable(b, jay.MaxUint16, jay.LenInt(jay.MaxUint16))
	assert.Equal(t, []byte{3, 255, 255, 0}, b)
	jay.WriteIntVariable(b, jay.MaxInt24, jay.LenInt(jay.MaxInt24))
	assert.Equal(t, []byte{3, 255, 255, 127}, b)
	b = make([]byte, 5)
	jay.WriteIntVariable(b, jay.MaxUint24, jay.LenInt(jay.MaxUint24))
	assert.Equal(t, []byte{4, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, jay.MaxInt32, jay.LenInt(jay.MaxInt32))
	assert.Equal(t, []byte{4, 255, 255, 255, 127}, b)
	b = make([]byte, 6)
	jay.WriteIntVariable(b, jay.MaxUint32, jay.LenInt(jay.MaxUint32))
	assert.Equal(t, []byte{5, 255, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, jay.MaxInt40, jay.LenInt(jay.MaxInt40))
	assert.Equal(t, []byte{5, 255, 255, 255, 255, 127}, b)
	b = make([]byte, 7)
	jay.WriteIntVariable(b, jay.MaxUint40, jay.LenInt(jay.MaxUint40))
	assert.Equal(t, []byte{6, 255, 255, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, jay.MaxInt48, jay.LenInt(jay.MaxInt48))
	assert.Equal(t, []byte{6, 255, 255, 255, 255, 255, 127}, b)
	b = make([]byte, 8)
	jay.WriteIntVariable(b, jay.MaxUint48, jay.LenInt(jay.MaxUint48))
	assert.Equal(t, []byte{7, 255, 255, 255, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, jay.MaxInt56, jay.LenInt(jay.MaxInt56))
	assert.Equal(t, []byte{7, 255, 255, 255, 255, 255, 255, 127}, b)
	b = make([]byte, 9)
	jay.WriteIntVariable(b, jay.MaxUint56, jay.LenInt(jay.MaxUint56))
	assert.Equal(t, []byte{8, 255, 255, 255, 255, 255, 255, 255, 0}, b)
	jay.WriteIntVariable(b, MaxInt64, jay.LenInt(MaxInt64))
	assert.Equal(t, []byte{8, 255, 255, 255, 255, 255, 255, 255, 127}, b)
}

func TestReadInt24(t *testing.T) {
	b := make([]byte, 3)
	// Increment each iteration by between 75 and 100.
	inc := rand.Intn(25) + 75
	for i := jay.MinInt24; i <= jay.MaxInt24; i += inc {
		jay.WriteInt24(b, i)
		require.Equal(t, i, jay.ReadInt24(b), b)
	}
}

/*
// ReadInt24 ...
func ReadInt24(b []byte) int {
	//if b[0]&_128 == 0 {

	value := int(b[0]) | int(b[1])<<8 | int(b[2])<<16
	return value | -(value & 0x00800000) // Apply sign extension if the highest bit is set
	//}

	//return | Neg24Mask int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16
	///}
	//return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16
}
*/

func TestReadInt40(t *testing.T) {
	b := make([]byte, 5)
	inc := jay.MaxUint32 - rand.Intn(150)
	for i := jay.MinInt40; i <= jay.MaxInt40; i += inc {
		jay.WriteInt40(b, i)
		require.Equal(t, i, jay.ReadInt40(b), b)
	}
}

func TestReadInt48(t *testing.T) {
	b := make([]byte, 6)
	inc := jay.MaxUint32 - rand.Intn(150)
	for i := jay.MinInt48; i <= jay.MaxInt48; i += inc {
		jay.WriteInt48(b, i)
		require.Equal(t, i, jay.ReadInt48(b), b)
	}
}

func TestReadInt56(t *testing.T) {
	b := make([]byte, 7)
	inc := jay.MaxUint48 - rand.Intn(150)
	for i := jay.MinInt56; i <= jay.MaxInt56; i += inc {
		jay.WriteInt56(b, i)
		require.Equal(t, i, jay.ReadInt56(b), b)
	}
}
