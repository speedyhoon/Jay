package jay

import (
	"time"
)

func ReadTime(b []byte) (t time.Time) {
	return time.Unix(ReadInt64(b), 0)
	//time.Second
}

// WriteTime writes 8 bytes to b with year, month, date, hour, minute & seconds precision in UTC location.
// All millisecond, microsecond, nanosecond and location is lost.
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
