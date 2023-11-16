package jay

const (
	MaxInt8  = 1<<7 - 1  // 127
	MinInt8  = -1 << 7   // -128
	MaxInt16 = 1<<15 - 1 // 32767
	MinInt16 = -1 << 15  // -32768
	MaxInt24 = 1<<23 - 1 // 8388607
	MinInt24 = -1 << 23  // -8388608
	MaxInt32 = 1<<31 - 1 // 2147483647
	MinInt32 = -1 << 31  // -2147483648
	MaxInt40 = 1<<39 - 1 // 549755813887
	MinInt40 = -1 << 39  // -549755813888
	MaxInt48 = 1<<47 - 1 // 140737488355327
	MinInt48 = -1 << 47  // -140737488355328
	MaxInt56 = 1<<55 - 1 // 36028797018963967
	MinInt56 = -1 << 55  // -36028797018963968
	MaxInt64 = 1<<63 - 1 // 9223372036854775807
	MinInt64 = -1 << 63  // -9223372036854775808
)

// ReadInt64 ...
func ReadInt64(b []byte) int64 {
	return int64(b[0]) | int64(b[1])<<_8 | int64(b[2])<<_16 | int64(b[3])<<_24 |
		int64(b[4])<<_32 | int64(b[5])<<_40 | int64(b[6])<<_48 | int64(b[7])<<_56
}

// ReadInt32 ...
func ReadInt32(b []byte) int32 {
	return int32(b[0]) | int32(b[1])<<_8 | int32(b[2])<<_16 | int32(b[3])<<_24
}

// ReadInt16 ...
func ReadInt16(b []byte) int16 {
	return int16(b[0]) | int16(b[1])<<_8
}

// WriteInt64 ...
func WriteInt64(b []byte, i int64) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48), byte(i>>_56)
}

// WriteInt32 ...
func WriteInt32(b []byte, i int32) {
	b[0], b[1], b[2], b[3] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24)
}

// WriteInt16 ...
func WriteInt16(b []byte, i int16) {
	b[0], b[1] = byte(i), byte(i>>_8)
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
		b[2] = byte(i >> _8)
	case 3:
		b[1] = byte(i)
		b[2] = byte(i >> _8)
		b[3] = byte(i >> _16)
	case 4:
		b[1] = byte(i)
		b[2] = byte(i >> _8)
		b[3] = byte(i >> _16)
		b[4] = byte(i >> _24)
	case 5:
		b[1] = byte(i)
		b[2] = byte(i >> _8)
		b[3] = byte(i >> _16)
		b[4] = byte(i >> _24)
		b[5] = byte(i >> _32)
	case 6:
		b[1] = byte(i)
		b[2] = byte(i >> _8)
		b[3] = byte(i >> _16)
		b[4] = byte(i >> _24)
		b[5] = byte(i >> _32)
		b[6] = byte(i >> _40)
	case 7:
		b[1] = byte(i)
		b[2] = byte(i >> _8)
		b[3] = byte(i >> _16)
		b[4] = byte(i >> _24)
		b[5] = byte(i >> _32)
		b[6] = byte(i >> _40)
		b[7] = byte(i >> _48)
	case 8:
		b[1] = byte(i)
		b[2] = byte(i >> _8)
		b[3] = byte(i >> _16)
		b[4] = byte(i >> _24)
		b[5] = byte(i >> _32)
		b[6] = byte(i >> _40)
		b[7] = byte(i >> _48)
		b[8] = byte(i >> _56)
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
