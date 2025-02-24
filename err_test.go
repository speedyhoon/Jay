package jay_test

import (
	"errors"
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrUnexpectedEOB(t *testing.T) {
	assert.Equal(t, errors.New("unexpected EOB"), jay.ErrUnexpectedEOB)
}
