package jay

const (
	MaxUint8  = 1<<8 - 1  // 255
	MaxUint16 = 1<<16 - 1 // 65535
	MaxUint24 = 1<<24 - 1 // 16777215
	MaxUint32 = 1<<32 - 1 // 4294967295
	MaxUint40 = 1<<40 - 1 // 1099511627775
	MaxUint48 = 1<<48 - 1 // 281474976710655
	MaxUint56 = 1<<56 - 1 // 72057594037927935
)

const (
	Neg24Mask = ^MaxUint24 // Negative integer masks.
	Neg40Mask = ^MaxUint40
	Neg48Mask = ^MaxUint48
	Neg56Mask = ^MaxUint56
)

// ReadInt24 ...
func ReadInt24(b []byte) int {
	// Check if the negative bit is on.
	if b[2]&_128 == _128 {
		return Neg24Mask | int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16
	}
	return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16
}

// ReadInt40 ...
func ReadInt40(b []byte) int {
	// Check if the negative bit is on.
	if b[4]&_128 == _128 {
		return Neg40Mask | int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24 | int(b[4])<<_32
	}
	return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24 | int(b[4])<<_32
}

// ReadInt48 ...
func ReadInt48(b []byte) int {
	// Check if the negative bit is on.
	if b[5]&_128 == _128 {
		return Neg48Mask | int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24 |
			int(b[4])<<_32 | int(b[5])<<_40
	}
	return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24 |
		int(b[4])<<_32 | int(b[5])<<_40
}

// ReadInt56 ...
func ReadInt56(b []byte) int {
	// Check if the negative bit is on.
	if b[6]&_128 == _128 {
		return Neg56Mask | int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24 |
			int(b[4])<<_32 | int(b[5])<<_40 | int(b[6])<<_48
	}

	return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24 |
		int(b[4])<<_32 | int(b[5])<<_40 | int(b[6])<<_48
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

// WriteIntVariable ...
func WriteIntVariable(b []byte, i int, length int) {
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
		return int(ReadInt16(b[1:])), 3
	case 3:
		return ReadInt24(b[1:]), 4
	case 4:
		return int(ReadInt32(b[1:])), 5
	case 5:
		return ReadInt40(b[1:]), 6
	case 6:
		return ReadInt48(b[1:]), 7
	case 7:
		return ReadInt56(b[1:]), 8
	case 8:
		return int(ReadInt64(b[1:])), 9
	}
	return
}

// WriteInt56 ...
func WriteInt56(b []byte, i int) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48)
}

// WriteInt48 ...
func WriteInt48(b []byte, i int) {
	b[0], b[1], b[2], b[3], b[4], b[5] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40)
}

// WriteInt40 ...
func WriteInt40(b []byte, i int) {
	b[0], b[1], b[2], b[3], b[4] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32)
}

// WriteInt24 ...
func WriteInt24(b []byte, i int) {
	b[0], b[1], b[2] = byte(i), byte(i>>_8), byte(i>>_16)
}
