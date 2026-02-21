package time_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

/*func BenchmarkReadTimeDate(b *testing.B) {
	for b.Loop() {
		timeTime = ReadTimeDate([]byte{240, 48, 77})
	}
}*/

func BenchmarkReadTime(b *testing.B) {
	for b.Loop() {
		timeTime = jay.ReadTime([]byte{240, 48, 77, 101, 225, 65, 87, 6})
	}
}

func BenchmarkReadTimeNano(b *testing.B) {
	for b.Loop() {
		timeTime = jay.ReadTimeNano([]byte{240, 48, 77, 101, 225, 65, 87, 6})
	}
}
