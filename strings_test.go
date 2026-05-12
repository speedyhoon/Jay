package jay_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/rando"
)

func TestRoundTripStringsArray(t *testing.T) {
	const size = 255

	var b []byte
	var list, actual [size]string
	t.Run("zero", func(t *testing.T) {
		jay.WriteStringsArray(b, 0, list[:])
		assert.NoError(t, jay.ReadStringsArrayErr(b, actual[:], 0))
		assert.Equal(t, list, actual)
	})

	result, seed := rando.StringsQtyLenSeed(size, size)
	expected := [size]string(result)
	b = make([]byte, jay.SizeStringsArray(expected[:], size))
	jay.WriteStringsArray(b, uint8(size), expected[:size])
	assert.NoErrorf(t, jay.ReadStringsArrayErr(b, actual[:], uint8(size)), "seed=%d", seed)
	assert.Equalf(t, expected, actual, "seed=%d", seed)
}
