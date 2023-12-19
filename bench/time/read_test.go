package time_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

/*func BenchmarkReadTimeDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timeTime = ReadTimeDate([]byte{240, 48, 77})
	}
}*/

func BenchmarkReadTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timeTime = jay.ReadTime([]byte{240, 48, 77, 101, 225, 65, 87, 6})
	}
}

func BenchmarkReadTimeNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timeTime = jay.ReadTimeNano([]byte{240, 48, 77, 101, 225, 65, 87, 6})
	}
}
