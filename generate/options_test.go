package generate

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_bytesRequired(t *testing.T) {
	var i, mx uint = 0, jay.MaxUint24 + 16
	var expected uint8
	for ; i < mx; i++ {
		switch i {
		case 1:
			expected = 1
		case jay.MaxUint8 + 1:
			expected = 2
		case jay.MaxUint16 + 1:
			expected = 3
		case jay.MaxUint24 + 1:
			expected = 4
		case jay.MaxUint32 + 1:
			expected = 5
		case jay.MaxUint40 + 1:
			expected = 6
		case jay.MaxUint48 + 1:
			expected = 7
		case jay.MaxUint56 + 1:
			expected = 8
		}

		require.Equalf(t, expected, bytesRequired(i), "input: %d, expected: %d", i, expected)
	}
}
