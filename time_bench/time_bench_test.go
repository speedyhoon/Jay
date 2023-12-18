package time_bench

import (
	"github.com/speedyhoon/jay"
	"testing"
	"time"
)

var timeTime time.Time

func BenchmarkReadTimeDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timeTime = ReadTimeDate([]byte{240, 48, 77})
	}
}

func BenchmarkWriteTimeDate(b *testing.B) {
	src := []byte{0, 0, 0}
	for i := 0; i < b.N; i++ {
		WriteTimeDate(src, timeTime)
	}
}

func BenchmarkWriteTimeDate2(b *testing.B) {
	src := []byte{0, 0, 0}
	for i := 0; i < b.N; i++ {
		WriteTimeDate2(src, timeTime)
	}
}

func BenchmarkReadTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timeTime = jay.ReadTime([]byte{240, 48, 77, 101, 225, 65, 87, 6})
	}
}

func BenchmarkWriteTime(b *testing.B) {
	src := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		jay.WriteTime(src, timeTime)
	}
}

func BenchmarkReadTimeNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timeTime /*, _*/ = ReadTimeNano([]byte{240, 48, 77, 101, 225, 65, 87, 6})
	}
}

func BenchmarkWriteTimeNano(b *testing.B) {
	src := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		_ = jay.WriteTimeNano(src, timeTime)
	}
}
