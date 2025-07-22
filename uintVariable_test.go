package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestWriteUint56(t *testing.T) {
	y := make([]byte, 7)
	jay.WriteUint56(y, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0}, y)

	jay.WriteUint56(y, math.MaxInt64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255}, y)
}

func TestWriteUint48(t *testing.T) {
	y := make([]byte, 6)
	jay.WriteUint48(y, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0}, y)

	jay.WriteUint48(y, math.MaxInt64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255, 255}, y)
}

func TestWriteUint40(t *testing.T) {
	y := make([]byte, 5)
	jay.WriteUint40(y, 0)
	assert.Equal(t, []byte{0, 0, 0, 0, 0}, y)

	jay.WriteUint40(y, math.MaxInt64)
	assert.Equal(t, []byte{255, 255, 255, 255, 255}, y)
}

func TestWriteUint24(t *testing.T) {
	y := make([]byte, 3)
	jay.WriteUint24(y, 0)
	assert.Equal(t, []byte{0, 0, 0}, y)

	jay.WriteUint24(y, math.MaxInt64)
	assert.Equal(t, []byte{255, 255, 255}, y)
}
