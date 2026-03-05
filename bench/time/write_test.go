package time_test

import (
	"testing"
	"time"

	"github.com/speedyhoon/jay"
)

var timeTime time.Time

/*func BenchmarkWriteTimeDate(b *testing.B) {
	src := []byte{0, 0, 0}
	for b.Loop() {
		WriteTimeDate(src, timeTime)
	}
}*/

/*func BenchmarkWriteTimeDate2(b *testing.B) {
	src := []byte{0, 0, 0}
	for b.Loop() {
		WriteTimeDate2(src, timeTime)
	}
}*/

func BenchmarkWriteTime(b *testing.B) {
	src := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for b.Loop() {
		jay.WriteTime(src, timeTime)
	}
}

func BenchmarkWriteTimeNano(b *testing.B) {
	src := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for b.Loop() {
		jay.WriteTimeNano(src, timeTime)
	}
}
