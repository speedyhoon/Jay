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
	if y[_2] >= _128 {
		return neg24Mask | int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16
	}
	return int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16
}

// ReadInt40 ...
func ReadInt40(y []byte) int {
	// Check if the negative bit is on.
	if y[_4] >= _128 {
		return neg40Mask | int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16 | int(y[_3])<<_24 | int(y[_4])<<_32
	}
	return int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16 | int(y[_3])<<_24 | int(y[_4])<<_32
}

// ReadInt48 ...
func ReadInt48(y []byte) int {
	// Check if the negative bit is on.
	if y[_5] >= _128 {
		return neg48Mask | int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16 | int(y[_3])<<_24 |
			int(y[_4])<<_32 | int(y[_5])<<_40
	}
	return int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16 | int(y[_3])<<_24 |
		int(y[_4])<<_32 | int(y[_5])<<_40
}

// ReadInt56 ...
func ReadInt56(y []byte) int {
	// Check if the negative bit is on.
	if y[_6] >= _128 {
		return neg56Mask | int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16 | int(y[_3])<<_24 |
			int(y[_4])<<_32 | int(y[_5])<<_40 | int(y[_6])<<_48
	}

	return int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16 | int(y[_3])<<_24 |
		int(y[_4])<<_32 | int(y[_5])<<_40 | int(y[_6])<<_48
}

// WriteIntVariable ...
func WriteIntVariable(y []byte, i int, length int) {
	y[_0] = byte(length)
	switch length {
	case _1:
		y[_1] = byte(i)
	case _2:
		y[_1], y[_2] = byte(i), byte(i>>_8)
	case _3:
		y[_1], y[_2], y[_3] = byte(i), byte(i>>_8), byte(i>>_16)
	case _4:
		y[_1], y[_2], y[_3], y[_4] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24)
	case _5:
		y[_1], y[_2], y[_3], y[_4], y[_5] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32)
	case _6:
		y[_1], y[_2], y[_3], y[_4], y[_5], y[_6] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40)
	case _7:
		y[_1], y[_2], y[_3], y[_4], y[_5], y[_6], y[_7] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48)
	case _8:
		y[_1], y[_2], y[_3], y[_4], y[_5], y[_6], y[_7], y[_8] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48), byte(i>>_56)
	}
}

// ReadInt ...
func ReadInt(y []byte) (i int, length int) {
	switch y[_0] {
	case _1:
		return int(int8(y[_1])), _2
	case _2:
		return int(ReadInt16(y[_1:_3])), _3
	case _3:
		return ReadInt24(y[_1:_4]), _4
	case _4:
		return int(ReadInt32(y[_1:_5])), _5
	case _5:
		return ReadInt40(y[_1:_6]), _6
	case _6:
		return ReadInt48(y[_1:_7]), _7
	case _7:
		return ReadInt56(y[_1:_8]), _8
	case _8:
		return int(ReadInt64(y[_1:9])), 9
	}
	return
}

// WriteInt56 ...
func WriteInt56(y []byte, i int) {
	y[_0], y[_1], y[_2], y[_3], y[_4], y[_5], y[_6] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48)
}

// WriteInt48 ...
func WriteInt48(y []byte, i int) {
	y[_0], y[_1], y[_2], y[_3], y[_4], y[_5] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40)
}

// WriteInt40 ...
func WriteInt40(y []byte, i int) {
	y[_0], y[_1], y[_2], y[_3], y[_4] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32)
}

// WriteInt24 ...
func WriteInt24(y []byte, i int) {
	y[_0], y[_1], y[_2] = byte(i), byte(i>>_8), byte(i>>_16)
}
