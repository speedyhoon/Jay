package jay_test

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/rando"
	"github.com/speedyhoon/utl/tf"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestComplex64(t *testing.T) {
	b := make([]byte, 8)
	var c complex64
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint(c), func(t *testing.T) {
			jay.WriteComplex64(b, c)
			assert.Equal(t, c, jay.ReadComplex64(b))
		})
		c = rando.Complex64()
	}
}

func TestComplex128(t *testing.T) {
	b := make([]byte, 16)
	var c complex128
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint(c), func(t *testing.T) {
			jay.WriteComplex128(b, c)
			assert.Equal(t, c, jay.ReadComplex128(b))
		})
		c = rando.Complex128()
	}
}

func TestRoundTripComplex64s(t *testing.T) {
	var b []byte
	var list []complex64
	t.Run("zero", func(t *testing.T) {
		jay.WriteComplex64s(b, list)
		assert.Equal(t, list, jay.ReadComplex64s(b, 0))
	})

	list = make([]complex64, math.MaxUint8)
	for i := range list {
		list[i] = rando.Complex64()
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*8)
			jay.WriteComplex64s(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadComplex64s(b, i))
		})
	}
}

func TestRoundTripComplex128s(t *testing.T) {
	var b []byte
	var list []complex128
	t.Run("zero", func(t *testing.T) {
		jay.WriteComplex128s(b, list)
		assert.Equal(t, list, jay.ReadComplex128s(b, 0))
	})

	list = make([]complex128, math.MaxUint8)
	for i := range list {
		list[i] = rando.Complex128()
	}

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*16)
			jay.WriteComplex128s(b, list[:i])
			assert.Equal(t, list[:i], jay.ReadComplex128s(b, i))
		})
	}
}
