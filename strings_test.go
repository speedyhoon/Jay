package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoundTripStringsArray(t *testing.T) {
	const size = 5

	var b []byte
	var list, actual [size]string
	t.Run("zero", func(t *testing.T) {
		jay.WriteStringsArray(b, 0, list[:])
		assert.NoError(t, jay.ReadStringsArrayErr(b, actual[:], 0))
		assert.Equal(t, list, actual)
	})

	expected := [size]string(rando.StringsN(size))
	b = make([]byte, jay.SizeStringsArray(expected[:], size))
	jay.WriteStringsArray(b, uint8(size), expected[:size])
	assert.NoError(t, jay.ReadStringsArrayErr(b, actual[:], uint8(size)))
	assert.Equal(t, expected, actual)
}
