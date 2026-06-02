package jay

const neg24Mask = ^(1<<24 - 1) // Negative integer masks.

// ReadInt24 ...
func ReadInt24(y []byte) int {
	// Check if the negative bit is on.
	if y[_2] >= _128 {
		return neg24Mask | int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16
	}
	return int(y[_0]) | int(y[_1])<<_8 | int(y[_2])<<_16
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
