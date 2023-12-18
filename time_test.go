package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now().UTC()
	b := make([]byte, 8)
	jay.WriteTime(b, now)
	assert.Equal(t, now.Format(time.RFC3339Nano), jay.ReadTime(b).Format(time.RFC3339Nano))
}

func TestTimeNil(t *testing.T) {
	zero := time.Time{}
	b := make([]byte, 8)
	jay.WriteTime(b, zero)
	assert.Equal(t, zero.Format(time.RFC3339Nano), jay.ReadTime(b).Format(time.RFC3339Nano))
}

func TestTimeNanoNow(t *testing.T) {

	now := time.Now().UTC()
	b := make([]byte, 8)
	jay.WriteTimeNano(b, now)
	assert.Equal(t, now.Format(time.RFC3339Nano), jay.ReadTimeNano(b).Format(time.RFC3339Nano))
}

func TestTimeNano98(t *testing.T) {
	now := time.Time{}
	b := make([]byte, 8)
	jay.WriteTimeNano(b, now)
	assert.Equal(t, now.Format(time.RFC3339Nano), jay.ReadTimeNano(b).Format(time.RFC3339Nano))
}

func TestWriteTimeNano3(t *testing.T) {
	b := make([]byte, 8)
	tests := map[string]time.Time{
		"now":   time.Now(),
		"empty": {},
		"zero":  time.Time{}.Add(0),
		"day1":  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		"1677":  time.Date(1677, time.September, 21, 00, 12, 44, 0, time.UTC),
		"min":   time.Unix(0, MinInt64),
		"1754":  time.Date(1754, time.August, 30, 22, 43, 41, 0, time.UTC),
		"1970":  time.Unix(0, 0),
		"max":   time.Unix(0, MaxInt64),
		"2262":  time.Date(2262, time.April, 11, 23, 47, 16, 0, time.UTC),
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			jay.WriteTimeNano(b, tt)
			assert.Equal(t, tt.UTC().Format(time.RFC3339Nano), jay.ReadTimeNano(b).Format(time.RFC3339Nano))
		})
	}
}

func TestWriteTime555555(t *testing.T) {
	b := make([]byte, 8)
	tests := map[string]time.Time{
		"now":   time.Now(),
		"empty": {},
		"zero":  time.Time{}.Add(0),
		"day1":  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		"1677":  time.Date(1677, time.September, 21, 00, 12, 44, 0, time.UTC),
		"min":   time.Unix(MinInt64, 0),
		"min2":  time.Unix(0, MinInt64),
		"min3":  time.Unix(MinInt64, MinInt64),
		"1754":  time.Date(1754, time.August, 30, 22, 43, 41, 0, time.UTC),
		"1970":  time.Unix(0, 0),
		"max":   time.Unix(MaxInt64, 0),
		"max2":  time.Unix(0, MaxInt64),
		"max3":  time.Unix(MaxInt64, MaxInt64),
		"2262":  time.Date(2262, time.April, 11, 23, 47, 16, 0, time.UTC),
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			jay.WriteTime(b, tt)
			assert.Equal(t, tt.UTC().Format(time.RFC3339), jay.ReadTime(b).Format(time.RFC3339))
		})
	}
}

func TestWriteTimeNano3Zero(t *testing.T) {
	input := time.Unix(0, MaxInt64-7)
	b := make([]byte, 8)

	t.Run("time zero",
		func(t *testing.T) {
			jay.WriteTimeNano(b, input)
			assert.Equal(t, time.Time{}.UTC().Format(time.RFC3339Nano), jay.ReadTimeNano(b).Format(time.RFC3339Nano))
		})
}
