package jay

import (
	"github.com/speedyhoon/rando"
	"github.com/speedyhoon/utl/tf"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRoundTripInt16s(t *testing.T) {
	var b []byte
	var list []int16
	t.Run("zero", func(t *testing.T) {
		WriteInt16s(b, list, 0)
		assert.Equal(t, list, ReadInt16s(b, 0))
	})

	list = rando.Int16sN(math.MaxUint8)

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*2)
			WriteInt16s(b, list[:i], i)
			assert.Equal(t, list[:i], ReadInt16s(b, i))
		})
	}
}

func TestRoundTripIntsX32(t *testing.T) {
	var b []byte
	var list []int
	t.Run("zero", func(t *testing.T) {
		WriteIntsX32(b, list)
		assert.Equal(t, list, ReadIntsX32(b, 0))
	})

	list = make([]int, math.MaxUint8)
	for i := range list {
		list[i] = int(rando.Int32())
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*4)
			WriteIntsX32(b, list[:i])
			assert.Equal(t, list[:i], ReadIntsX32(b, i))
		})
	}
}

func TestRoundTripIntsX64(t *testing.T) {
	var b []byte
	var list []int
	t.Run("zero", func(t *testing.T) {
		WriteIntsX64(b, list)
		assert.Equal(t, list, ReadIntsX64(b, 0))
	})

	list = make([]int, math.MaxUint8)
	for i := range list {
		list[i] = int(rando.Int64())
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*8)
			WriteIntsX64(b, list[:i])
			assert.Equal(t, list[:i], ReadIntsX64(b, i))
		})
	}
}

func TestRoundTripInt64s(t *testing.T) {
	var b []byte
	var list []int64
	t.Run("zero", func(t *testing.T) {
		WriteInt64s(b, list)
		assert.Equal(t, list, ReadInt64s(b, 0))
	})

	list = make([]int64, math.MaxUint8)
	for i := range list {
		list[i] = rando.Int64()
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*8)
			WriteInt64s(b, list[:i])
			assert.Equal(t, list[:i], ReadInt64s(b, i))
		})
	}
}

func TestRoundTripInt32s(t *testing.T) {
	var b []byte
	var list []int32
	t.Run("zero", func(t *testing.T) {
		WriteInt32s(b, list)
		assert.Equal(t, list, ReadInt32s(b, 0))
	})

	list = make([]int32, math.MaxUint8)
	for i := range list {
		list[i] = rando.Int32()
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*4)
			WriteInt32s(b, list[:i])
			assert.Equal(t, list[:i], ReadInt32s(b, i))
		})
	}
}
