package generate_test

import (
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/jay/generate"
	"github.com/speedyhoon/jay/generate/tests"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var _ generate.Option

func TestBytes(t *testing.T) {
	//var opt generate.Option
	//assert.NoError(t, opt.ProcessWrite(nil, "./tests/boolJay.go", "./tests/bools.go"))

	var cpy tests.OnlyBools

	t.Run("empty struct", func(t *testing.T) {
		ob := tests.OnlyBools{}
		src := ob.MarshalJ()
		assert.Equal(t, []byte{0}, src)

		assert.Equal(t, jay.ErrUnexpectedEOB, cpy.UnmarshalJ(nil))
		assert.Equal(t, jay.ErrUnexpectedEOB, cpy.UnmarshalJ([]byte{}))
		assert.NoError(t, cpy.UnmarshalJ(src))
	})

	t.Run("random values", func(t *testing.T) {
		ob := tests.OnlyBools{
			Hidden:      randBool(),
			Deactivated: randBool(),
			Selected:    randBool(),
			Modified:    randBool(),
			Updated:     randBool(),
		}
		src := ob.MarshalJ()
		assert.NoError(t, cpy.UnmarshalJ(src))
		assert.Equal(t, ob.Hidden, cpy.Hidden)
		assert.Equal(t, ob.Deactivated, cpy.Deactivated)
		assert.Equal(t, ob.Selected, cpy.Selected)
		assert.Equal(t, ob.Modified, cpy.Modified)
		assert.Equal(t, ob.Updated, cpy.Updated)
	})

}

func randBool() bool {
	return rand.Intn(1) == 1
}
