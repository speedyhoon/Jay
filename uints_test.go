package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/rando"
	"github.com/speedyhoon/utl/tf"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRoundTripUint16s(t *testing.T) {
	var b []byte
	var list []uint16
	t.Run("zero", func(t *testing.T) {
		jay.WriteUint16s(b, list, 0)
		assert.Equal(t, list, jay.ReadUint16s(b, 0))
	})

	list = rando.Uint16sN(math.MaxInt8)

	for i := 1; i <= math.MaxInt8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*2)
			jay.WriteUint16s(b, list[:i], i)
			assert.Equal(t, list[:i], jay.ReadUint16s(b, i))
		})
	}
}

func TestRoundTripUintsX32(t *testing.T) {
	var b []byte
	var list []uint
	t.Run("zero", func(t *testing.T) {
		jay.WriteUintsX32(b, list)
		assert.Equal(t, list, jay.ReadUintsX32(b, 0))
	})

	list = make([]uint, math.MaxUint8)
	for i := range list {
		list[i] = uint(rando.Uint32())
	}

	for i := 1; i <= math.MaxInt8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*4)
			jay.WriteUintsX32(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadUintsX32(b, i))
		})
	}
}

func TestRoundTripUintsX64(t *testing.T) {
	var b []byte
	var list []uint
	t.Run("zero", func(t *testing.T) {
		jay.WriteUintsX64(b, list)
		assert.Equal(t, list, jay.ReadUintsX64(b, 0))
	})

	list = make([]uint, math.MaxUint8)
	for i := range list {
		list[i] = uint(rando.Uint64())
	}

	for i := 1; i <= math.MaxInt8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*8)
			jay.WriteUintsX64(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadUintsX64(b, i))
		})
	}
}

func TestRoundTripUint64s(t *testing.T) {
	var b []byte
	var list []uint64
	t.Run("zero", func(t *testing.T) {
		jay.WriteUint64s(b, list)
		assert.Equal(t, list, jay.ReadUint64s(b, 0))
	})

	list = make([]uint64, math.MaxUint8)
	for i := range list {
		list[i] = rando.Uint64()
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*8)
			jay.WriteUint64s(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadUint64s(b, i))
		})
	}
}

func TestRoundTripUint32s(t *testing.T) {
	var b []byte
	var list []uint32
	t.Run("zero", func(t *testing.T) {
		jay.WriteUint32s(b, list)
		assert.Equal(t, list, jay.ReadUint32s(b, 0))
	})

	list = make([]uint32, math.MaxUint8)
	for i := range list {
		list[i] = rando.Uint32()
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*4)
			jay.WriteUint32s(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadUint32s(b, i))
		})
	}
}
