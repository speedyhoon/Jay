package generate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_byteSize(t *testing.T) {
	tests := []struct {
		input tagSize
		want  uint
	}{
		{input: 0, want: 0},
		{input: 1, want: 1},
		{input: 255, want: 1},
		{input: 256, want: 2},
		{input: 65535, want: 2},
		{input: 65536, want: 3},
		{input: 16777215, want: 3},
		{input: 16777216, want: 4},
		{input: 4294967295, want: 4},
		{input: 4294967296, want: 5},
		{input: 1099511627775, want: 5},
		{input: 1099511627776, want: 6},
		{input: 281474976710655, want: 6},
		{input: 281474976710656, want: 7},
		{input: 72057594037927935, want: 7},
		{input: 72057594037927936, want: 8},
		{input: 18446744073709551614, want: 8},
		{input: 18446744073709551615, want: 8},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			assert.Equalf(t, tt.want, byteSize(tt.input), "byteSize(%v)", tt.input)
		})
	}
}

func Test_isShortRequiredTag(t *testing.T) {
	tests := []struct {
		tag  string
		want bool
	}{
		{tag: "undefined", want: false},
		{tag: "unmatched", want: false},
		{tag: "unnamed", want: false},
		{tag: "needed", want: false},
		{tag: "required", want: true},
		{tag: "require", want: true},
		{tag: "requir", want: true},
		{tag: "requi", want: true},
		{tag: "requ", want: true},
		{tag: "req", want: true},
		{tag: "re", want: true},
		{tag: "r", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			assert.Equalf(t, tt.want, isShortRequiredTag(tt.tag), "isShortRequiredTag(%v)", tt.tag)
		})
	}
}
