package generate_test

import (
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/jay/generate"
	"github.com/speedyhoon/jay/generate/tests"
	"github.com/speedyhoon/jay/rando"
	"github.com/stretchr/testify/assert"
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
			Hidden:      rando.Bool(),
			Deactivated: rando.Bool(),
			Selected:    rando.Bool(),
			Modified:    rando.Bool(),
			Updated:     rando.Bool(),
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
