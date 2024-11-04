package generate_test

import (
	"github.com/speedyhoon/jay/generate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExportedErrs(t *testing.T) {
	assert.Equal(t, "no filename or source provided", generate.ErrNoSource.Error())
	assert.Equal(t, "no exported struct fields found", generate.ErrNoneExported.Error())
}
