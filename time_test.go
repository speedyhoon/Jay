package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	b := make([]byte, 8)
	jay.WriteTime(b, now)
	assert.Equal(t, now.Format(time.RFC3339), jay.ReadTime(b).Format(time.RFC3339))
}

func TestTimeNano(t *testing.T) {
	now := time.Now()
	b := make([]byte, 15)
	assert.NoError(t, jay.WriteTimeNano(b, now))

	r, err := jay.ReadTimeNano(b)
	assert.NoError(t, err)
	assert.Equal(t, now.Format(time.RFC3339Nano), r.Format(time.RFC3339Nano))
}
