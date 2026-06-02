package jay

const (
	neg40Mask int = ^(1<<40 - 1)
	neg48Mask int = ^(1<<48 - 1)
	neg56Mask int = ^(1<<56 - 1)
)

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

// ReadInt for 64-bit systems.
func ReadInt(y []byte) (i, length int) {
	switch y[_0] {
	case _1:
		return int(int8(y[_1])), _2 // Convert to int8 to ensure negative integers are handled.
	case _2:
		return int(ReadInt16(y[_1:_3])), _3
	case _3:
		return ReadInt24(y[_1:_4]), _4
	case _4:
		return ReadIntX32(y[_1:_5]), _5
	case _5:
		return ReadInt40(y[_1:_6]), _6
	case _6:
		return ReadInt48(y[_1:_7]), _7
	case _7:
		return ReadInt56(y[_1:_8]), _8
	case _8:
		return ReadIntX64(y[_1:9]), 9
	}
	return
}
