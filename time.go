package jay

import (
	"time"
)

func ReadTime(b []byte) (t time.Time) {
	return time.Unix(ReadInt64(b), 0)
	//time.Second
}

// WriteTime writes 8 bytes to b with year, month, date, hour, minute & seconds precision in UTC location.
// All millisecond, microsecond, nanosecond and location are lost.
func WriteTime(b []byte, t time.Time) {
	WriteInt64(b, t.Unix())
}

func ReadTimeNano(b []byte) (t time.Time, err error) {
	err = t.UnmarshalBinary(b) //time.Unm(ReadInt64(b), 0)
	return
}

// WriteTimeNano writes 15 bytes to b with year, month, date, hour, minute, seconds, millisecond,
// microsecond & nanosecond precision with any valid location.
func WriteTimeNano(b []byte, t time.Time) (err error) {
	var s []byte
	s, err = t.MarshalBinary()
	if err == nil {
		copy(b[:15], s)
	}
	return
}

// WriteTimeDate uses 3 bytes to represent year, month and day.
// The supported year range is 21,844 AD to 21,844 BC.
func WriteTimeDate(b []byte, t time.Time) {
	year, month, day := t.Date()
	month-- // Months start at 1 instead of 0, so decrementing increases the date range.

	var negativeYearFlag byte
	if year < 0 {
		year *= -12
		negativeYearFlag = 128
	} else {
		year *= 12
	}

	// Year and month are combined to increase the date range available. Otherwise, bits are wasted representing months. 2^4-1 = 15, not 12.
	year += int(month)

	b[0] = negativeYearFlag | byte(year>>11)
	b[1] = byte(year >> 3)
	b[2] = byte(year<<5 | day)
}

// WriteTimeDate2 uses 3 bytes to represent year, month and day.
// The supported year range is 21,844 AD to 21,844 BC.
func WriteTimeDate2(b []byte, t time.Time) {
	year := t.Year()
	//month-- // Months start at 1 instead of 0, so decrementing increases the date range.

	var negativeYearFlag byte
	if year < 0 {
		year *= -366
		negativeYearFlag = 128
	} else {
		year *= 366
	}

	// Year and month are combined to increase the date range available. Otherwise, bits are wasted representing months. 2^4-1 = 15, not 12.
	//year += int(month)
	yearDay := year + t.YearDay() - 1

	b[0], b[1], b[2] = negativeYearFlag|byte(yearDay>>_16), byte(yearDay>>_8), byte(yearDay)

	//b[0] = negativeYearFlag | byte(year>>11)
	//b[1] = byte(year >> 3)
	//b[2] = byte(year<<5 | day)
}

func ReadTimeDate(b []byte) time.Time {

	//if b[2]&_128 == _128 {
	//	return Neg24Mask | int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16
	//}
	//return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16
	//
	yearDays := int(b[0]&127)<<16 | int(b[1])<<8 | int(b[2])

	// If the negative year flag is present, multiply the year to be negative.
	negativeYearFlag := int(b[0]&128)/-64 + 1

	return time.Date(
		yearDays/366*negativeYearFlag,
		1, //time.Month(yearDays%12+1),
		yearDays%366+1,
		0, 0, 0, 0, time.UTC,
	)
} /*

func ReadTimeDate(b []byte) time.Time {
	yearMonths := int(b[0]&127)<<11 | int(b[1])<<3 | int(b[2])>>5

	// If the negative year flag is present, multiply the year to be negative.
	negativeYearFlag := int(b[0]&128)/-64 + 1

	return time.Date(
		yearMonths/12*negativeYearFlag,
		time.Month(yearMonths%12+1),
		int(b[2]&31),
		0, 0, 0, 0, time.UTC,
	)
}*/
