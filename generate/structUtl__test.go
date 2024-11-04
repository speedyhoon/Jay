package generate_test

import (
	"github.com/speedyhoon/jay/generate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_exportedConstants(t *testing.T) {
	assert.Equal(t, "-", generate.IgnoreFlag)
	assert.Equal(t, "j", generate.StructTagName)
}
