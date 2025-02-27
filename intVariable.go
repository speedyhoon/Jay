package jay

const (
	neg24Mask = ^(1<<24 - 1) // Negative integer masks.
	neg40Mask = ^(1<<40 - 1)
	neg48Mask = ^(1<<48 - 1)
	neg56Mask = ^(1<<56 - 1)
)

// ReadInt24 ...
func ReadInt24(y []byte) int {
	// Check if the negative bit is on.
	if y[2]&_128 == _128 {
		return neg24Mask | int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16
	}
	return int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16
}

// ReadInt40 ...
func ReadInt40(y []byte) int {
	// Check if the negative bit is on.
	if y[4]&_128 == _128 {
		return neg40Mask | int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16 | int(y[3])<<_24 | int(y[4])<<_32
	}
	return int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16 | int(y[3])<<_24 | int(y[4])<<_32
}

// ReadInt48 ...
func ReadInt48(y []byte) int {
	// Check if the negative bit is on.
	if y[5]&_128 == _128 {
		return neg48Mask | int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16 | int(y[3])<<_24 |
			int(y[4])<<_32 | int(y[5])<<_40
	}
	return int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16 | int(y[3])<<_24 |
		int(y[4])<<_32 | int(y[5])<<_40
}

// ReadInt56 ...
func ReadInt56(y []byte) int {
	// Check if the negative bit is on.
	if y[6]&_128 == _128 {
		return neg56Mask | int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16 | int(y[3])<<_24 |
			int(y[4])<<_32 | int(y[5])<<_40 | int(y[6])<<_48
	}

	return int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16 | int(y[3])<<_24 |
		int(y[4])<<_32 | int(y[5])<<_40 | int(y[6])<<_48
}

// WriteIntVariable ...
func WriteIntVariable(y []byte, i int, length int) {
	y[0] = byte(length)
	switch length {
	case 1:
		y[1] = byte(i)
	case 2:
		y[1] = byte(i)
		y[2] = byte(i >> _8)
	case 3:
		y[1] = byte(i)
		y[2] = byte(i >> _8)
		y[3] = byte(i >> _16)
	case 4:
		y[1] = byte(i)
		y[2] = byte(i >> _8)
		y[3] = byte(i >> _16)
		y[4] = byte(i >> _24)
	case 5:
		y[1] = byte(i)
		y[2] = byte(i >> _8)
		y[3] = byte(i >> _16)
		y[4] = byte(i >> _24)
		y[5] = byte(i >> _32)
	case 6:
		y[1] = byte(i)
		y[2] = byte(i >> _8)
		y[3] = byte(i >> _16)
		y[4] = byte(i >> _24)
		y[5] = byte(i >> _32)
		y[6] = byte(i >> _40)
	case 7:
		y[1] = byte(i)
		y[2] = byte(i >> _8)
		y[3] = byte(i >> _16)
		y[4] = byte(i >> _24)
		y[5] = byte(i >> _32)
		y[6] = byte(i >> _40)
		y[7] = byte(i >> _48)
	case 8:
		y[1] = byte(i)
		y[2] = byte(i >> _8)
		y[3] = byte(i >> _16)
		y[4] = byte(i >> _24)
		y[5] = byte(i >> _32)
		y[6] = byte(i >> _40)
		y[7] = byte(i >> _48)
		y[8] = byte(i >> _56)
	}
}

// ReadInt ...
func ReadInt(y []byte) (i int, length int) {
	switch y[0] {
	case 1:
		return int(int8(y[1])), 2
	case 2:
		return int(ReadInt16(y[1:3])), 3
	case 3:
		return ReadInt24(y[1:4]), 4
	case 4:
		return int(ReadInt32(y[1:5])), 5
	case 5:
		return ReadInt40(y[1:6]), 6
	case 6:
		return ReadInt48(y[1:7]), 7
	case 7:
		return ReadInt56(y[1:8]), 8
	case 8:
		return int(ReadInt64(y[1:9])), 9
	}
	return
}

// WriteInt56 ...
func WriteInt56(y []byte, i int) {
	y[0], y[1], y[2], y[3], y[4], y[5], y[6] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48)
}

// WriteInt48 ...
func WriteInt48(y []byte, i int) {
	y[0], y[1], y[2], y[3], y[4], y[5] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40)
}

// WriteInt40 ...
func WriteInt40(y []byte, i int) {
	y[0], y[1], y[2], y[3], y[4] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32)
}

// WriteInt24 ...
func WriteInt24(y []byte, i int) {
	y[0], y[1], y[2] = byte(i), byte(i>>_8), byte(i>>_16)
}
