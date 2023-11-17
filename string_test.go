package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	const str = "octopus camouflage"
	const length = len(str)
	y := make([]byte, length+1)

	jay.String(y, str, length)
	assert.Equal(t, []byte("\022octopus camouflage"), y)

	s, l, ok := jay.ReadString(z)
	assert.Equal(t, str, s)
	assert.Equal(t, length+1, l)
	assert.True(t, ok)
}

func TestStringPtr(t *testing.T) {
	const str = "octopus camouflage"
	const length = len(str)
	y := make([]byte, length+1)

	jay.String(y, str, length)
	assert.Equal(t, []byte("\022octopus camouflage"), y)

	l, ok := jay.ReadStringPtr(z, &str1)
	assert.Equal(t, str, str1)
	assert.Equal(t, length+1, l)
	assert.True(t, ok)
}

func TestStrings(t *testing.T) {
	list := []string{"octopus", "camouflage"}
	length := 1 + len(list) + len(list[0]) + len(list[1])
	b := make([]byte, length)

	jay.Strings(b, list)
	assert.Equal(t, []byte("\002\007octopus\012camouflage"), b)

	s, l, ok := jay.ReadStrings(b)
	assert.Equal(t, list, s)
	assert.Equal(t, length, l)
	assert.True(t, ok)
}
