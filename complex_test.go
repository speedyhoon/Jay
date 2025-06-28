package jay_test

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComplex64(t *testing.T) {
	b := make([]byte, 8)
	var c complex64
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint(c), func(t *testing.T) {
			jay.WriteComplex64(b, c)
			assert.Equal(t, c, jay.ReadComplex64(b))
		})
		c = rando.Complex64()
	}
}

func TestComplex128(t *testing.T) {
	b := make([]byte, 16)
	var c complex128
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint(c), func(t *testing.T) {
			jay.WriteComplex128(b, c)
			assert.Equal(t, c, jay.ReadComplex128(b))
		})
		c = rando.Complex128()
	}
}
