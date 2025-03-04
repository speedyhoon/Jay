package jay_test

import (
	"fmt"
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrings(t *testing.T) {
	list := []string{"octopus", "camouflage"}
	qty := len(list)
	length := qty + len(list[0]) + len(list[1])
	b := make([]byte, length)

	jay.WriteStrings(b[:qty], b[qty:], list)
	assert.Equal(t, []byte{7, 10, 'o', 'c', 't', 'o', 'p', 'u', 's', 'c', 'a', 'm', 'o', 'u', 'f', 'l', 'a', 'g', 'e'}, b)

	//s, l, ok := jay.ReadStrings(b)
	//assert.Equal(t, list, s)
	//assert.Equal(t, length, l)
	//assert.True(t, ok)
}

func TestSizeStrings8(t *testing.T) {
	tests := []struct {
		wantTotal int
		wantQty   int
		wantSizes []byte
		s         []string
	}{
		{wantTotal: 20, wantQty: 3, wantSizes: []byte{3, 3, 10}, s: []string{"one", "two", "hello 😄"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d] total:%d size:%d", i, tt.wantTotal, tt.wantQty), func(t *testing.T) {
			gotTotal /*, gotQty, gotSizes*/ := jay.SizeStrings8(tt.s)
			assert.Equalf(t, tt.wantTotal, gotTotal, "SizeStrings8(%v)", tt.s)
			//assert.Equalf(t, tt.wantQty, gotQty, "SizeStrings8(%v)", tt.s)
			//assert.Equalf(t, tt.wantSizes, gotSizes, "SizeStrings8(%v)", tt.s)
		})
	}
}

func TestSizeStrings16(t *testing.T) {
	tests := []struct {
		wantTotal int
		wantQty   int
		wantSizes []byte
		s         []string
	}{
		{wantTotal: 24, wantQty: 3, wantSizes: []byte{3, 0, 3, 0, 10, 0}, s: []string{"one", "two", "hello 😄"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d] total:%d size:%d", i, tt.wantTotal, tt.wantQty), func(t *testing.T) {
			gotTotal /*, gotQty, gotSizes*/ := jay.SizeStrings16(tt.s)
			assert.Equalf(t, tt.wantTotal, gotTotal, "SizeStrings16(%v)", tt.s)
			//assert.Equalf(t, tt.wantQty, gotQty, "SizeStrings16(%v)", tt.s)
			//assert.Equalf(t, tt.wantSizes, gotSizes, "SizeStrings16(%v)", tt.s)
		})
	}
}

/*func TestSizeStrings24(t *testing.T) {
	tests := []struct {
		wantTotal int
		wantQty   int
		wantSizes []byte
		s         []string
	}{
		{wantTotal: 16, wantQty: 3, wantSizes: []byte{3, 0, 0, 3, 0, 0, 10, 0, 0}, s: []string{"one", "two", "hello 😄"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d] total:%d size:%d", i, tt.wantTotal, tt.wantQty), func(t *testing.T) {
			gotTotal, gotQty, gotSizes := jay.SizeStrings24(tt.s)
			assert.Equalf(t, tt.wantTotal, gotTotal, "SizeStrings24(%v)", tt.s)
			assert.Equalf(t, tt.wantQty, gotQty, "SizeStrings24(%v)", tt.s)
			assert.Equalf(t, tt.wantSizes, gotSizes, "SizeStrings24(%v)", tt.s)
		})
	}
}*/
