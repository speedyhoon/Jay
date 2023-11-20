package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolsSliceIndex(t *testing.T) {
	tests := map[uint]uint{
		0: 0,
		1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0,
		9: 8, 10: 8, 11: 8, 12: 8, 13: 8, 14: 8, 15: 8, 16: 8,
		17: 16, 18: 16, 19: 16, 20: 16, 21: 16, 22: 16, 23: 16, 24: 16,
		25: 24, 26: 24, 27: 24, 28: 24, 29: 24, 30: 24, 31: 24, 32: 24,
		33: 32, 34: 32, 35: 32, 36: 32, 37: 32, 38: 32, 39: 32, 40: 32,
		41: 40, 42: 40, 43: 40, 44: 40, 45: 40, 46: 40, 47: 40, 48: 40,
		49: 48, 50: 48, 51: 48, 52: 48, 53: 48, 54: 48, 55: 48, 56: 48,
		57: 56, 58: 56, 59: 56, 60: 56, 61: 56, 62: 56, 63: 56, 64: 56,
		65: 64, 66: 64, 67: 64, 68: 64, 69: 64, 70: 64, 71: 64, 72: 64,
	}
	for input, expected := range tests {
		assert.Equal(t, expected, jay.BoolsSliceIndex(input), "input: %d", input)
	}
}
