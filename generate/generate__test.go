package generate_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/speedyhoon/jay/generate"
)

func TestExportedErrs(t *testing.T) {
	assert.Equal(t, "no filename or source provided", generate.ErrNoSource.Error())
	assert.Equal(t, "no exported struct fields found", generate.ErrNoneExported.Error())
}
