package jay

const (
	MaxInt8  = 1<<7 - 1  // 127
	MinInt8  = -1 << 7   // -128
	MaxInt16 = 1<<15 - 1 // 32767
	MinInt16 = -1 << 15  // -32768
	MaxInt24 = 1<<23 - 1 // 32767
	MinInt24 = -1 << 23  // -32768
	MaxInt32 = 1<<31 - 1 // 2147483647
	MinInt32 = -1 << 31  // -2147483648
	MaxInt40 = 1<<39 - 1 // 2147483647
	MinInt40 = -1 << 39  // -2147483648
	MaxInt48 = 1<<47 - 1 // 2147483647
	MinInt48 = -1 << 47  // -2147483648
	MaxInt56 = 1<<55 - 1 // 2147483647
	MinInt56 = -1 << 55  // -2147483648
)

// ReadInt64 ...
func ReadInt64(b []byte) int64 {
	return int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
		int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
}

// ReadInt32 ...
func ReadInt32(b []byte) int32 {
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

// ReadInt16 ...
func ReadInt16(b []byte) int16 {
	return int16(b[0]) | int16(b[1])<<8
}

// WriteInt64 ...
func WriteInt64(b []byte, i int64) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(i), byte(i>>8), byte(i>>16), byte(i>>24), byte(i>>32), byte(i>>40), byte(i>>48), byte(i>>56)
}

// WriteInt32 ...
func WriteInt32(b []byte, i int32) {
	b[0], b[1], b[2], b[3] = byte(i), byte(i>>8), byte(i>>16), byte(i>>24)
}

// WriteInt16 ...
func WriteInt16(b []byte, i int16) {
	b[0], b[1] = byte(i), byte(i>>8)
}

// LenInt ...
func LenInt(i int) int {
	if i == 0 {
		return 0
	}

	if i > MaxInt56 || i < MinInt56 {
		return 8
	}

	if i > MaxInt48 || i < MinInt48 {
		return 7
	}

	if i > MaxInt40 || i < MinInt40 {
		return 6
	}

	if i > MaxInt32 || i < MinInt32 {
		return 5
	}

	if i > MaxInt24 || i < MinInt24 {
		return 4
	}

	if i > MaxInt16 || i < MinInt16 {
		return 3
	}

	if i > MaxInt8 || i < MinInt8 {
		return 2
	}

	return 1
}

// WriteInt ...
func WriteInt(b []byte, i int, length int) {
	b[0] = byte(length)
	switch length {
	case 1:
		b[1] = byte(i)
	case 2:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
	case 3:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
	case 4:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
	case 5:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		b[5] = byte(i >> 32)
	case 6:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		b[5] = byte(i >> 32)
		b[6] = byte(i >> 40)
	case 7:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		b[5] = byte(i >> 32)
		b[6] = byte(i >> 40)
		b[7] = byte(i >> 48)
	case 8:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		b[5] = byte(i >> 32)
		b[6] = byte(i >> 40)
		b[7] = byte(i >> 48)
		b[8] = byte(i >> 56)
	}
}

// ReadInt ...
func ReadInt(b []byte) (i int, length int) {
	switch b[0] {
	case 1:
		return int(int8(b[1])), 2
	case 2:
		return int(ReadInt16(b)), 3
	case 3:
		return int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24, 4
	case 4:
		return int(ReadInt32(b)), 5
	case 5:
		return int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24 | int(b[4])<<32, 6
	case 6:
		return int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24 |
			int(b[4])<<32 | int(b[5])<<40, 7
	case 7:
		return int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24 |
			int(b[4])<<32 | int(b[5])<<40 | int(b[6])<<48, 8
	case 8:
		return int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24 |
			int(b[4])<<32 | int(b[5])<<40 | int(b[6])<<48 | int(b[7])<<56, 9
	}
	return
}
