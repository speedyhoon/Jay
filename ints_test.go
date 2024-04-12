package jay

import (
	"github.com/speedyhoon/rando"
	"github.com/speedyhoon/utl/tf"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRoundTripInt16s(t *testing.T) {
	var b []byte
	var list []int16
	t.Run("zero", func(t *testing.T) {
		WriteInt16s(b, list, 0)
		assert.Equal(t, list, ReadInt16s(b, 0))
	})

	list = rando.Int16sN(math.MaxUint8)

	for i := 1; i <= math.MaxUint8; i++ {
		tf.Run(t, i, func(t *testing.T) {
			b = make([]byte, i*2)
			WriteInt16s(b, list[:i], i)
			assert.Equal(t, list[:i], ReadInt16s(b, i))
		})
	}
}
