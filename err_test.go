package jay_test

import (
	"errors"
	"github.com/go-openapi/testify/v2/assert"
	"github.com/speedyhoon/jay"
	"testing"
)

func TestErrUnexpectedEOB(t *testing.T) {
	assert.Equal(t, errors.New("unexpected EOB"), jay.ErrUnexpectedEOB)
}
