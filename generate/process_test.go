package generate

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	zero := make([]Struct, 0)
	one := []Struct{{name: "one"}}
	two := []Struct{{name: "one"}, {name: "two"}}

	tests := []struct {
		s     []Struct
		index int
		want  []Struct
	}{
		{s: nil, index: -1, want: nil},
		{s: nil, index: 0, want: nil},
		{s: nil, index: 1, want: nil},

		{s: zero, index: -1, want: zero},
		{s: zero, index: 0, want: zero},
		{s: zero, index: 1, want: zero},

		{s: one, index: -1, want: one},
		{s: one, index: 0, want: zero},
		{s: one, index: 1, want: one},

		{s: two, index: -1, want: two},
		{s: two, index: 0, want: []Struct{{name: "two"}}},
		{s: two, index: 1, want: one},
		{s: two, index: 2, want: two},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tt.want, Remove(tt.s, tt.index))
		})
	}
}
