package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoundTripStringsArray(t *testing.T) {
	var b []byte
	var list [5]string
	t.Run("zero", func(t *testing.T) {
		jay.WriteStringsArray(b, 0, list[:])
		var actual [5]string
		assert.NoError(t, jay.ReadStringsArrayErr(b, actual[:], 0))
		assert.Equal(t, list, actual)
	})

	const size = 5
	expected := [size]string(rando.StringsN(size))
	b = make([]byte, jay.SizeStringsArray(expected[:], size))
	var actual [size]string
	jay.WriteStringsArray(b, uint8(size), expected[:size])
	assert.NoError(t, jay.ReadStringsArrayErr(b, actual[:], uint8(size)))
	assert.Equal(t, expected, actual)
}
