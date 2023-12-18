package jay

import (
	"time"
)

func ReadTime(b []byte) (t time.Time) {
	return time.Unix(ReadInt64(b), 0).UTC()
}

// WriteTime writes 8 bytes to b with year, month, date, hour, minute & seconds precision in UTC location.
// All millisecond, microsecond, nanosecond and location are lost.
func WriteTime(b []byte, t time.Time) {
	WriteInt64(b, t.Unix())
}

const timeZero = MaxInt64 - 7 // 9223372036854775800

func ReadTimeNano(b []byte) (t time.Time) {
	if i := ReadInt64(b); i != timeZero {
		return time.Unix(0, i).UTC()
	}

	return
}

// WriteTimeNano writes 8 bytes to b with year, month, date, hour, minute & seconds precision in UTC location.
// All millisecond, microsecond and nanosecond are lost.
func WriteTimeNano(b []byte, t time.Time) {
	if t == (time.Time{}) {
		b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = 248, 255, 255, 255, 255, 255, 255, 127 // WriteInt64(b, timeZero)
		return
	}

	WriteInt64(b, t.UnixNano())
}
