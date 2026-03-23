package genjay_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/speedyhoon/jay/genjay"
)

func TestExportedErrs(t *testing.T) {
	assert.Equal(t, "no filename or source provided", genjay.ErrNoSource.Error())
	assert.Equal(t, "no exported struct fields found", genjay.ErrNoneExported.Error())
}
