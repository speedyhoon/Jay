package jay

import "time"

const maxUint8 = 255

type Usize interface {
	~[]bool |
		~[]complex128 |
		~[]complex64 |
		~[]float32 |
		~[]float64 |
		~[]int |
		~[]int16 |
		~[]int32 |
		~[]int64 |
		~[]int8 |
		~[]string |
		~[]time.Duration |
		~[]time.Time |
		~[]uint |
		~[]uint16 |
		~[]uint32 |
		~[]uint64 |
		~[]uint8 |
		~string
}

func Len8[T Usize](v T) int {
	return min(len(v), maxUint8)
}
