package jay

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

func TestFloat32(t *testing.T) {
	b := make([]byte, 4)
	for i := 0; i < 1000; i++ {
		f := rand.Float32() * float32(math.Pow(10, float64(rand.Intn(9)+1)))
		name := strconv.FormatFloat(float64(f), 'f', -1, 32)
		t.Run(name, func(t *testing.T) {
			WriteFloat32(b, f)
			assert.Equal(t, f, ReadFloat32(b))
		})
	}
}

func TestFloat32MaxSmallest(t *testing.T) {
	b := make([]byte, 4)
	t.Run("max", func(t *testing.T) {
		WriteFloat32(b, math.MaxFloat32)
		assert.Equal(t, float32(math.MaxFloat32), ReadFloat32(b))
	})
	t.Run("smallest", func(t *testing.T) {
		WriteFloat32(b, math.SmallestNonzeroFloat32)
		assert.Equal(t, float32(math.SmallestNonzeroFloat32), ReadFloat32(b))
	})
}

func TestFloat64(t *testing.T) {
	b := make([]byte, 8)
	for i := 0; i < 1000; i++ {
		f := rand.Float64() * math.Pow(10, float64(rand.Intn(9)+1))
		name := strconv.FormatFloat(f, 'f', -1, 64)
		t.Run(name, func(t *testing.T) {
			WriteFloat64(b, f)
			assert.Equal(t, f, ReadFloat64(b))
		})
	}
}

func TestFloat64MaxSmallest(t *testing.T) {
	b := make([]byte, 8)
	t.Run("max", func(t *testing.T) {
		WriteFloat64(b, math.MaxFloat64)
		assert.Equal(t, math.MaxFloat64, ReadFloat64(b))
	})
	t.Run("smallest", func(t *testing.T) {
		WriteFloat64(b, math.SmallestNonzeroFloat64)
		assert.Equal(t, math.SmallestNonzeroFloat64, ReadFloat64(b))
	})
}
