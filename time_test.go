package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
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

func TestTimeDate(t *testing.T) {
	//const maxValue = 7978842 // 21845-04-01
	const maxValue = 8371246  // 22919-09-11
	const minValue = -8370628 // 22919-09-11
	_, output := make([]byte, 3), make([]byte, 3)
	//var expectedDate time.Time // Date(-1, time.January, 1, 0, 0, 0, 0, time.UTC)
	//expectedDate := time.Date(-1, time.December, 31, 0, 0, 0, 0, time.UTC)
	//expectedDate := time.Date(-1, time.December, 1, 0, 0, 0, 0, time.UTC)
	//expectedDate := time.Date(21845, time.April, 30, 0, 0, 0, 0, time.UTC)

	//expectedDate := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
	expectedDate := time.Date(-22918, time.January, 1, 0, 0, 0, 0, time.UTC)
	var maxYear, minYear int

	//for i := 0; i <= maxValue; i++ {
	for i := 0; i >= -maxValue; i-- {
		//if i == 365 {
		//	println()
		//}
		jay.WriteTimeDate2(output, expectedDate)

		date := jay.ReadTimeDate(output)

		//jay.WriteInt24(input, i)
		//input[0], input[2] = input[2], input[0]

		//if i == 31 {
		//	println("hello")
		//}
		//date := jay.ReadTimeDate(input)
		require.Equalf(t, expectedDate, date, "index: %d, bytes: %v", i, output)

		//jay.WriteTimeDate(output, date)
		//require.Equal(t, input, output)

		expectedDate = expectedDate.AddDate(0, 0, -1)
		if date.Year() > maxYear {
			maxYear = date.Year()
		}
		if date.Year() < minYear {
			minYear = date.Year()
		}
	}
	log.Println("max", maxYear, "min", minYear)
}

/*
func TestTimeDate(t *testing.T) {
	input, _ := make([]byte, 3), make([]byte, 3)
	//var expectedDate time.Time // Date(-1, time.January, 1, 0, 0, 0, 0, time.UTC)
	//expectedDate := time.Date(-1, time.December, 31, 0, 0, 0, 0, time.UTC)
	expectedDate := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
	var maxYear, minYear int

	for i := 0; i < jay.MaxUint24; i++ {
		jay.WriteInt24(input, i)
		input[0], input[2] = input[2], input[0]

		if i == 31 {
			println("hello")
		}
		date := jay.ReadTimeDate(input)
		require.Equalf(t, expectedDate, date, "index: %d", i)

		//jay.WriteTimeDate(output, date)
		//require.Equal(t, input, output)

		expectedDate = expectedDate.AddDate(0, 0, 1)
		if date.Year() > maxYear {
			maxYear = date.Year()
		}
		if date.Year() < minYear {
			minYear = date.Year()
		}
	}
	log.Println("max", maxYear, "min", minYear)
}
*/
