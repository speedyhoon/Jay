package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrings(t *testing.T) {
	list := []string{"octopus", "camouflage"}
	qty := len(list)
	length := qty + len(list[0]) + len(list[1])
	b := make([]byte, length)

	jay.WriteStrings(b[:qty], b[qty:], list)
	assert.Equal(t, []byte{7, 10, 'o', 'c', 't', 'o', 'p', 'u', 's', 'c', 'a', 'm', 'o', 'u', 'f', 'l', 'a', 'g', 'e'}, b)

	//s, l, ok := jay.ReadStrings(b)
	//assert.Equal(t, list, s)
	//assert.Equal(t, length, l)
	//assert.True(t, ok)
}
