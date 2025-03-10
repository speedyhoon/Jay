package jay

import "time"

func ReadTime(y []byte) (t time.Time) {
	return time.Unix(ReadInt64(y), 0).UTC()
}

// WriteTime writes 8 bytes to b with year, month, date, hour, minute and seconds precision in UTC location.
// All millisecond, microsecond, nanosecond and location are lost.
func WriteTime(y []byte, t time.Time) {
	WriteInt64(y, t.Unix())
}

func ReadTimes(y []byte, length int) (t []time.Time) {
	if length == 0 {
		return
	}

	t = make([]time.Time, length)
	for i := 0; i < length; i++ {
		t[i] = ReadTime(y[i*_8 : i*_8+_8])
	}
	return
}

func WriteTimes(y []byte, t []time.Time) {
	for i := range t {
		WriteInt64(y[i*_8:i*_8+_8], t[i].Unix())
	}
}

// timeZero represents the undefined value for time.Time.
// Zero (0) is not used to prevent errors when time.Unix(0) (1970-01-01) are converted to 0001-01-01.
// The maximum and minimum values math.MaxInt64 and math.MinInt64 are also not used to prevent collisions
// with the earliest and last time.Time representable.
// Seven nanoseconds are deducted from math.MaxInt64 to provide sufficient buffer against the maximum.
const timeZero = 1<<63 - 8 // (math.MaxInt64-7) 9223372036854775800

func ReadTimeNano(y []byte) (t time.Time) {
	if i := ReadInt64(y); i != timeZero {
		return time.Unix(0, i).UTC()
	}

	return
}

// WriteTimeNano writes 8 bytes to b with year, month, date, hour, minute and seconds precision in UTC location.
// All millisecond, microsecond and nanosecond are lost.
func WriteTimeNano(y []byte, t time.Time) {
	if t == (time.Time{}) {
		y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7] = 248, 255, 255, 255, 255, 255, 255, 127 // WriteInt64(y, timeZero)
		return
	}

	WriteInt64(y, t.UnixNano())
}

// ReadDurations ...
func ReadDurations(y []byte, length int) (t []time.Duration) {
	if length == 0 {
		return
	}

	t = make([]time.Duration, length)
	for i := 0; i < length; i++ {
		t[i] = time.Duration(ReadInt64(y[i*_8 : i*_8+_8]))
	}
	return
}

// WriteDurations ...
func WriteDurations(y []byte, t []time.Duration) {
	for i, dur := range t {
		WriteInt64(y[i*_8:i*_8+_8], int64(dur))
	}
}
