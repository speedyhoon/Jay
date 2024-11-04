package generate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_tagSymbol(t *testing.T) {
	assert.Equal(t, '`', tagSymbol)
}

func Test_unwrapTagValue(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "none", input: "", want: ""},
		{name: "empty", input: "``", want: ""},
		{name: "single grave", input: "```", want: "`"},
		{name: "double grave", input: "````", want: "``"},
		{name: "triple grave", input: "`````", want: "```"},
		{name: "text", input: "hello", want: "hello"},
		{name: "code", input: "(b, x.Char != nil, x.Num != nil, x.Rip != nil, x.Deleted, x.Modified)", want: "(b, x.Char != nil, x.Num != nil, x.Rip != nil, x.Deleted, x.Modified)"},
		{name: "code wrapped", input: "`(b, x.Char != nil, x.Num != nil, x.Rip != nil, x.Deleted, x.Modified)`", want: "(b, x.Char != nil, x.Num != nil, x.Rip != nil, x.Deleted, x.Modified)"},
		{name: "unwrap", input: "`castle`", want: "castle"},
		{name: "tag value", input: "`j:abc`", want: "j:abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, unwrapTagValue(tt.input), "unwrapTagValue(%v)", tt.input)
		})
	}
}
