package jay_test

import (
	"github.com/speedyhoon/jay"
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
		jay.WriteInt16s(b, list, 0)
		assert.Equal(t, list, jay.ReadInt16s(b, 0))
	})

	list = rando.Int16sN(math.MaxUint8)

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*2)
			jay.WriteInt16s(b, list[:i], i)
			assert.Equal(t, list[:i], jay.ReadInt16s(b, i))
		})
	}
}

func TestRoundTripIntsX32(t *testing.T) {
	var b []byte
	var list []int
	t.Run("zero", func(t *testing.T) {
		jay.WriteIntsX32(b, list)
		assert.Equal(t, list, jay.ReadIntsX32(b, 0))
	})

	list = make([]int, math.MaxUint8)
	for i := range list {
		list[i] = int(rando.Int32())
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*4)
			jay.WriteIntsX32(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadIntsX32(b, i))
		})
	}
}

func TestRoundTripIntsX64(t *testing.T) {
	var b []byte
	var list []int
	t.Run("zero", func(t *testing.T) {
		jay.WriteIntsX64(b, list)
		assert.Equal(t, list, jay.ReadIntsX64(b, 0))
	})

	list = rando.IntsN(math.MaxUint8)

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*8)
			jay.WriteIntsX64(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadIntsX64(b, i))
		})
	}
}

func TestRoundTripInt64s(t *testing.T) {
	var b []byte
	var list []int64
	t.Run("zero", func(t *testing.T) {
		jay.WriteInt64s(b, list)
		assert.Equal(t, list, jay.ReadInt64s(b, 0))
	})

	list = rando.Int64sN(math.MaxUint8)

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*8)
			jay.WriteInt64s(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadInt64s(b, i))
		})
	}
}

func TestRoundTripInt32s(t *testing.T) {
	var b []byte
	var list []int32
	t.Run("zero", func(t *testing.T) {
		jay.WriteInt32s(b, list)
		assert.Equal(t, list, jay.ReadInt32s(b, 0))
	})

	list = rando.Int32sN(math.MaxUint8)

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*4)
			jay.WriteInt32s(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadInt32s(b, i))
		})
	}
}

func TestRoundTripInt8s(t *testing.T) {
	var b []byte
	var list []int8
	t.Run("zero", func(t *testing.T) {
		jay.WriteInt8s(b, list)
		assert.Equal(t, list, jay.ReadInt8s(b, 0))
	})

	list = rando.Int8sN(math.MaxUint8)

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*4)
			jay.WriteInt8s(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadInt8s(b, i))
		})
	}
}
