package generate_test

import (
	"github.com/speedyhoon/jay/generate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Constants_BitSize(t *testing.T) {
	assert.Equal(t, generate.MaxSize(0), generate.BitAuto, "BitAuto = %d, want: 0", generate.BitAuto)
	assert.Equal(t, generate.MaxSize(32), generate.Bit32, "Bit32 = %d, want: 32", generate.Bit32)
	assert.Equal(t, generate.MaxSize(64), generate.Bit64, "Bit64 = %d, want: 64", generate.Bit64)
}

func Test_Constants_FieldTags(t *testing.T) {
	assert.Equal(t, "-", generate.IgnoreFlag)
	assert.Equal(t, "j", generate.StructTagName)
}
