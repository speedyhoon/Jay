package jay_test

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteBools(t *testing.T) {
	tests := []struct {
		y    []byte
		a    []bool
		want []byte
	}{
		{y: nil, a: nil, want: nil},
		{y: []byte{0}, a: nil, want: []byte{0}},
		{y: []byte{0}, a: []bool{false}, want: []byte{0}},
		{y: []byte{0}, a: []bool{true}, want: []byte{128}},
		{y: []byte{0}, a: []bool{true, false}, want: []byte{128}},
		{y: []byte{0}, a: []bool{true, true}, want: []byte{192}},
		{y: []byte{0}, a: []bool{true, true, false}, want: []byte{192}},
		{y: []byte{0}, a: []bool{true, true, true}, want: []byte{224}},
		{y: []byte{0}, a: []bool{true, true, true, true}, want: []byte{240}},
		{y: []byte{0}, a: []bool{true, true, true, true, true, true, true, true}, want: []byte{255}},
		{y: []byte{0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true,
		}, want: []byte{255, 128}},
		{y: []byte{0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}, want: []byte{255, 255}},
		{y: []byte{0, 0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true,
		}, want: []byte{255, 255, 128}},
		{y: []byte{0, 0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}, want: []byte{255, 255, 255}},
		{y: []byte{0, 0, 0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true,
		}, want: []byte{255, 255, 255, 128}},
		{y: []byte{0, 0, 0, 0}, a: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}, want: []byte{255, 255, 255, 255}},

		{y: []byte{0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false,
		}, want: []byte{0, 0}},
		{y: []byte{0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}, want: []byte{0, 0}},
		{y: []byte{0, 0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false,
		}, want: []byte{0, 0, 0}},
		{y: []byte{0, 0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}, want: []byte{0, 0, 0}},
		{y: []byte{0, 0, 0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false,
		}, want: []byte{0, 0, 0, 0}},
		{y: []byte{0, 0, 0, 0}, a: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}, want: []byte{0, 0, 0, 0}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test: %d", i+1), func(t *testing.T) {
			jay.WriteBools(tt.y, tt.a, len(tt.a))
			assert.Equal(t, tt.want, tt.y)
		})
	}
}

func TestReadBools(t *testing.T) {
	tests := []struct {
		y    []byte
		want []bool
	}{
		{y: nil, want: nil},
		{y: []byte{0}, want: nil},
		{y: []byte{0}, want: []bool{false}},
		{y: []byte{128}, want: []bool{true}},
		{y: []byte{128}, want: []bool{true, false}},
		{y: []byte{192}, want: []bool{true, true}},
		{y: []byte{192}, want: []bool{true, true, false}},
		{y: []byte{224}, want: []bool{true, true, true}},
		{y: []byte{240}, want: []bool{true, true, true, true}},
		{y: []byte{255}, want: []bool{true, true, true, true, true, true, true, true}},
		{y: []byte{255, 128}, want: []bool{
			true, true, true, true, true, true, true, true,
			true,
		}},
		{y: []byte{255, 255}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}},
		{y: []byte{255, 255, 128}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true,
		}},
		{y: []byte{255, 255, 255}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}},
		{y: []byte{255, 255, 255, 128}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true,
		}},
		{y: []byte{255, 255, 255, 255}, want: []bool{
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
			true, true, true, true, true, true, true, true,
		}},

		{y: []byte{0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false,
		}},
		{y: []byte{0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}},
		{y: []byte{0, 0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false,
		}},
		{y: []byte{0, 0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}},
		{y: []byte{0, 0, 0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false,
		}},
		{y: []byte{0, 0, 0, 0}, want: []bool{
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, false, false,
		}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test: %d", i+1), func(t *testing.T) {
			assert.Equal(t, tt.want, jay.ReadBools(tt.y, len(tt.want)))
		})
	}
}
