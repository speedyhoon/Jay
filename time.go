package jay

import "time"

func ReadTime(b []byte) (t time.Time) {
	return time.Unix(ReadInt64(b), 0).UTC()
}

// WriteTime writes 8 bytes to b with year, month, date, hour, minute & seconds precision in UTC location.
// All millisecond, microsecond, nanosecond and location are lost.
func WriteTime(b []byte, t time.Time) {
	WriteInt64(b, t.Unix())
}

// timeZero represents the undefined value for time.Time.
// Zero (0) is not used to prevent errors when time.Unix(0) (1970-01-01) are converted to 0001-01-01.
// The maximum and minimum values math.MaxInt64 and math.MinInt64 are also not used to prevent collisions
// with the earliest and last time.Time representable.
// Seven nanoseconds are deducted from math.MaxInt64 to provide sufficient buffer against the maximum.
const timeZero = 1<<63 - 8 // (math.MaxInt64-7) 9223372036854775800

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

// ReadDuration ...
func ReadDuration(b []byte) time.Duration {
	return time.Duration(b[0]) | time.Duration(b[1])<<_8 | time.Duration(b[2])<<_16 | time.Duration(b[3])<<_24 |
		time.Duration(b[4])<<_32 | time.Duration(b[5])<<_40 | time.Duration(b[6])<<_48 | time.Duration(b[7])<<_56
}

// WriteDuration ...
func WriteDuration(b []byte, t time.Duration) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(t), byte(t>>_8), byte(t>>_16), byte(t>>_24), byte(t>>_32), byte(t>>_40), byte(t>>_48), byte(t>>_56)
}
