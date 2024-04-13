package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteUint64(t *testing.T) {
	y := make([]byte, 8)
	jay.WriteUint64(y, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, y)

	jay.WriteUint64(y, MaxInt64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 127}, y)
	jay.WriteUint64(y, MaxUint64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 255}, y)
}

func TestWriteUintX64(t *testing.T) {
	y := make([]byte, 8)
	jay.WriteUintX64(y, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, y)

	jay.WriteUintX64(y, MaxInt64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 127}, y)

	jay.WriteUintX64(y, MaxUint64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 255}, y)
}

func TestWriteUintX32(t *testing.T) {
	y := make([]byte, 4)
	jay.WriteUintX32(y, 0)
	assert.Equal(t, []byte{0, 0, 0, 0}, y)
	jay.WriteUintX32(y, MaxInt32)
	assert.Equal(t, []byte{255, 255, 255, 127}, y)
	jay.WriteUintX32(y, MaxUint64)
	assert.Equal(t, []byte{255, 255, 255, 255}, y)
}

func TestWriteUint32(t *testing.T) {
	y := make([]byte, 4)
	jay.WriteUint32(y, 0)
	assert.Equal(t, []byte{0, 0, 0, 0}, y)
	jay.WriteUint32(y, MaxInt32)
	assert.Equal(t, []byte{255, 255, 255, 127}, y)
	jay.WriteUint32(y, MaxUint32)
	assert.Equal(t, []byte{255, 255, 255, 255}, y)
}

func TestWriteUint16(t *testing.T) {
	y := make([]byte, 2)
	jay.WriteUint16(y, 0)
	assert.Equal(t, []byte{0, 0}, y)
	jay.WriteUint16(y, MaxInt16)
	assert.Equal(t, []byte{255, 127}, y)
	jay.WriteUint16(y, MaxUint16)
	assert.Equal(t, []byte{255, 255}, y)
}

func TestReadUint16(t *testing.T) {
	assert.Equal(t, uint16(0), jay.ReadUint16([]byte{0, 0}))
	assert.Equal(t, uint16(247), jay.ReadUint16([]byte{247, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint16(MaxUint16), jay.ReadUint16([]byte{255, 255, 255, 255, 255, 255, 255}))
}

func TestReadUint64(t *testing.T) {
	assert.Equal(t, uint64(0), jay.ReadUint64([]byte{0, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint64(247), jay.ReadUint64([]byte{247, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint64(MaxUint64), jay.ReadUint64([]byte{255, 255, 255, 255, 255, 255, 255, 255}))
}

func TestReadUintX64(t *testing.T) {
	assert.Equal(t, uint(0), jay.ReadUintX64([]byte{0, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint(247), jay.ReadUintX64([]byte{247, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint(MaxUint64), jay.ReadUintX64([]byte{255, 255, 255, 255, 255, 255, 255, 255}))
}

func TestReadUintX32(t *testing.T) {
	assert.Equal(t, uint(0), jay.ReadUintX32([]byte{0, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint(247), jay.ReadUintX32([]byte{247, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint(MaxUint32), jay.ReadUintX32([]byte{255, 255, 255, 255, 255, 255, 255, 255}))
}

func TestReadUint32(t *testing.T) {
	assert.Equal(t, uint32(0), jay.ReadUint32([]byte{0, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint32(247), jay.ReadUint32([]byte{247, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, uint32(MaxUint32), jay.ReadUint32([]byte{255, 255, 255, 255, 255, 255, 255, 255}))
}
