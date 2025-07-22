package jay

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, _0, 0)
	assert.Equal(t, _1, 1)
	assert.Equal(t, _2, 2)
	assert.Equal(t, _3, 3)
	assert.Equal(t, _4, 4)
	assert.Equal(t, _5, 5)
	assert.Equal(t, _6, 6)
	assert.Equal(t, _7, 7)
	assert.Equal(t, _8, 8)
	assert.Equal(t, _16, 16)
	assert.Equal(t, _24, 24)
	assert.Equal(t, _32, 32)
	assert.Equal(t, _40, 40)
	assert.Equal(t, _48, 48)
	assert.Equal(t, _56, 56)
	assert.Equal(t, _64, 64)
	assert.Equal(t, _128, 128)
	assert.Equal(t, maxUint8, 255)
}
