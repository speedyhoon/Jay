package generate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bytesRequired(t *testing.T) {
	for input, expected := range map[uint]uint8{
		0:             0,
		1:             1,
		maxUint8 - 1:  1,
		maxUint8:      1,
		maxUint8 + 1:  2,
		maxUint16 - 1: 2,
		maxUint16:     2,
		maxUint16 + 1: 3,
		maxUint24 - 1: 3,
		maxUint24:     3,
		maxUint24 + 1: 4,
		maxUint32 - 1: 4,
		maxUint32:     4,
		maxUint32 + 1: 5,
		maxUint40 - 1: 5,
		maxUint40:     5,
		maxUint40 + 1: 6,
		maxUint48 - 1: 6,
		maxUint48:     6,
		maxUint48 + 1: 7,
		maxUint56 - 1: 7,
		maxUint56:     7,
	} {
		output := bytesRequired(input)
		assert.Equalf(t, expected, output, "input: %d, expected: %d, output: %d", input, expected, output)
	}
}
