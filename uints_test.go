package jay

import (
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
		WriteUint16s(b, list, 0)
		assert.Equal(t, list, ReadUint16s(b, 0))
	})

	list = rando.Uint16sN(math.MaxInt8)

	for i := 1; i <= math.MaxInt8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*2)
			WriteUint16s(b, list[:i], i)
			assert.Equal(t, list[:i], ReadUint16s(b, i))
		})
	}
}

func TestRoundTripUintsArch32(t *testing.T) {
	var b []byte
	var list []uint
	t.Run("zero", func(t *testing.T) {
		WriteUintsX32(b, list)
		assert.Equal(t, list, ReadUintsX32(b, 0))
	})

	list = make([]uint, math.MaxUint8)
	for i := range list {
		list[i] = uint(rando.Uint32())
	}

	for i := 1; i <= math.MaxInt8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*4)
			WriteUintsX32(b, list[:i])
			assert.Equal(t, list[:i], ReadUintsX32(b, i))
		})
	}
}

func TestRoundTripUintsArch64(t *testing.T) {
	var b []byte
	var list []uint
	t.Run("zero", func(t *testing.T) {
		WriteUintsX64(b, list)
		assert.Equal(t, list, ReadUintsX64(b, 0))
	})

	list = make([]uint, math.MaxUint8)
	for i := range list {
		list[i] = uint(rando.Uint64())
	}

	for i := 1; i <= math.MaxInt8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*8)
			WriteUintsX64(b, list[:i])
			assert.Equal(t, list[:i], ReadUintsX64(b, i))
		})
	}
}
