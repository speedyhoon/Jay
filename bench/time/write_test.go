package time_test

import (
	"github.com/speedyhoon/jay"
	"testing"
	"time"
)

var timeTime time.Time

/*func BenchmarkWriteTimeDate(b *testing.B) {
	src := []byte{0, 0, 0}
	for i := 0; i < b.N; i++ {
		WriteTimeDate(src, timeTime)
	}
}*/

/*func BenchmarkWriteTimeDate2(b *testing.B) {
	src := []byte{0, 0, 0}
	for i := 0; i < b.N; i++ {
		WriteTimeDate2(src, timeTime)
	}
}*/

func BenchmarkWriteTime(b *testing.B) {
	src := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		jay.WriteTime(src, timeTime)
	}
}

func BenchmarkWriteTimeNano(b *testing.B) {
	src := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		jay.WriteTimeNano(src, timeTime)
	}
}
