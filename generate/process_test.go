package generate

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	zero := make([]structTyp, 0)
	one := []structTyp{{name: "one"}}
	two := []structTyp{{name: "one"}, {name: "two"}}

	tests := []struct {
		s     []structTyp
		index int
		want  []structTyp
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
		{s: two, index: 0, want: []structTyp{{name: "two"}}},
		{s: two, index: 1, want: one},
		{s: two, index: 2, want: two},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tt.want, Remove(tt.s, tt.index))
		})
	}
}
