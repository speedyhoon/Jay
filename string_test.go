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

func TestStringsSize8(t *testing.T) {
	tests := []struct {
		wantTotal int
		wantQty   int
		wantSizes []byte
		s         []string
	}{
		{wantTotal: 16, wantQty: 3, wantSizes: []byte{3, 3, 10}, s: []string{"one", "two", "hello 😄"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d] total:%d size:%d", i, tt.wantTotal, tt.wantQty), func(t *testing.T) {
			gotTotal, gotQty, gotSizes := jay.StringsSize8(tt.s)
			assert.Equalf(t, tt.wantTotal, gotTotal, "StringsSize8(%v)", tt.s)
			assert.Equalf(t, tt.wantQty, gotQty, "StringsSize8(%v)", tt.s)
			assert.Equalf(t, tt.wantSizes, gotSizes, "StringsSize8(%v)", tt.s)
		})
	}
}

func TestStringsSize16(t *testing.T) {
	tests := []struct {
		wantTotal int
		wantQty   int
		wantSizes []byte
		s         []string
	}{
		{wantTotal: 16, wantQty: 3, wantSizes: []byte{3, 0, 3, 0, 10, 0}, s: []string{"one", "two", "hello 😄"}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d] total:%d size:%d", i, tt.wantTotal, tt.wantQty), func(t *testing.T) {
			gotTotal, gotQty, gotSizes := jay.StringsSize16(tt.s)
			assert.Equalf(t, tt.wantTotal, gotTotal, "StringsSize16(%v)", tt.s)
			assert.Equalf(t, tt.wantQty, gotQty, "StringsSize16(%v)", tt.s)
			assert.Equalf(t, tt.wantSizes, gotSizes, "StringsSize16(%v)", tt.s)
		})
	}
}

func TestStringsSize24(t *testing.T) {
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
			gotTotal, gotQty, gotSizes := jay.StringsSize24(tt.s)
			assert.Equalf(t, tt.wantTotal, gotTotal, "StringsSize24(%v)", tt.s)
			assert.Equalf(t, tt.wantQty, gotQty, "StringsSize24(%v)", tt.s)
			assert.Equalf(t, tt.wantSizes, gotSizes, "StringsSize24(%v)", tt.s)
		})
	}
}
