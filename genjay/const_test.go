package genjay_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/speedyhoon/jay/genjay"
)

func Test_Constants_BitSize(t *testing.T) {
	assert.Equal(t, genjay.MaxSize(0), genjay.BitAuto, "BitAuto = %d, want: 0", genjay.BitAuto)
	assert.Equal(t, genjay.MaxSize(32), genjay.Bit32, "Bit32 = %d, want: 32", genjay.Bit32)
	assert.Equal(t, genjay.MaxSize(64), genjay.Bit64, "Bit64 = %d, want: 64", genjay.Bit64)
}

func Test_Constants_FieldTags(t *testing.T) {
	assert.Equal(t, "-", genjay.IgnoreFlag)
	assert.Equal(t, "j", genjay.StructTagName)
}
