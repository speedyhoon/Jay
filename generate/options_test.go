package generate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_bytesRequired(t *testing.T) {
	var i, mx uint = 0, MaxUint24 + 16
	var expected uint8
	for ; i < mx; i++ {
		switch i {
		case 1:
			expected = 1
		case MaxUint8 + 1:
			expected = 2
		case MaxUint16 + 1:
			expected = 3
		case MaxUint24 + 1:
			expected = 4
		case MaxUint32 + 1:
			expected = 5
		case MaxUint40 + 1:
			expected = 6
		case MaxUint48 + 1:
			expected = 7
		case MaxUint56 + 1:
			expected = 8
		}

		output := bytesRequired(i)
		require.Equalf(t, expected, bytesRequired(i), "input: %d, expected: %d, output: %d", i, expected, output)
	}
}

func TestLogBaseX(t *testing.T) {
	tests := []struct {
		base float64
		x    float64
		want float64
	}{
		{base: 3, x: 3, want: 1},
		{base: 3, x: 9, want: 2.0000000000000004},
		{base: 3, x: 27, want: 3.0000000000000004},
		{base: 4, x: 4, want: 1},
		{base: 4, x: 16, want: 2},
		{base: 4, x: 64, want: 3},
		{base: 4, x: 256, want: 4},
		{base: 4, x: 1024, want: 5},
		{base: 4, x: 4096, want: 6},
		{base: 4, x: 16384, want: 7},
		{base: 10, x: 10, want: 1},
		{base: 10, x: 100, want: 2},
		{base: 10, x: 1000, want: 2.9999999999999996},
		{base: 10, x: 10000, want: 4},
		{base: 256, x: 256, want: 1},
		{base: 256, x: 256 * 256, want: 2},
		{base: 256, x: 256 * 256 * 256, want: 3},
		{base: 256, x: 256 * 256 * 256 * 256, want: 4},
	}
	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%f, %f", tt.base, tt.x),
			func(t *testing.T) {
				assert.Equalf(t, tt.want, LogBaseX(tt.base, tt.x), "LogBaseX(%v, %v)", tt.base, tt.x)
			})
	}
}
