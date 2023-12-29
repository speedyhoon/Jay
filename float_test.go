package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func TestFloat32(t *testing.T) {
	b := make([]byte, 4)
	for i := 0; i < 1000; i++ {
		f := rando.Float32()
		t.Run(fmtFloat32(f), func(t *testing.T) {
			jay.WriteFloat32(b, f)
			assert.Equal(t, f, jay.ReadFloat32(b))
		})
	}
}

func TestFloat32MaxSmallest(t *testing.T) {
	b := make([]byte, 4)
	t.Run("max", func(t *testing.T) {
		jay.WriteFloat32(b, math.MaxFloat32)
		assert.Equal(t, float32(math.MaxFloat32), jay.ReadFloat32(b))
	})
	t.Run("smallestNonzeroFloat32", func(t *testing.T) {
		jay.WriteFloat32(b, math.SmallestNonzeroFloat32)
		assert.Equal(t, float32(math.SmallestNonzeroFloat32), jay.ReadFloat32(b))
	})
}

func TestFloat64(t *testing.T) {
	b := make([]byte, 8)
	for i := 0; i < 1000; i++ {
		f := rando.Float64()
		t.Run(fmtFloat64(f), func(t *testing.T) {
			jay.WriteFloat64(b, f)
			assert.Equal(t, f, jay.ReadFloat64(b))
		})
	}
}

func TestFloat64MaxSmallest(t *testing.T) {
	b := make([]byte, 8)
	t.Run("max", func(t *testing.T) {
		jay.WriteFloat64(b, math.MaxFloat64)
		assert.Equal(t, math.MaxFloat64, jay.ReadFloat64(b))
	})
	t.Run("smallestNonzeroFloat64", func(t *testing.T) {
		jay.WriteFloat64(b, math.SmallestNonzeroFloat64)
		assert.Equal(t, math.SmallestNonzeroFloat64, jay.ReadFloat64(b))
	})
}

func fmtFloat32(f float32) string {
	return fmtFloat64(float64(f))
}

func fmtFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 32)
}
