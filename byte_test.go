package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytes(t *testing.T) {
	src := []byte("octopus camouflage")
	length := len(src)
	y := make([]byte, length+1)

	jay.WriteBytes(y, src, length)
	assert.Equal(t, []byte("\022octopus camouflage"), y)

	z := []byte("\022octopus camouflage")
	s, size, ok := jay.ReadBytes(z)
	assert.Equal(t, src, s)
	assert.Equal(t, size, length+1)
	assert.True(t, ok)
}
