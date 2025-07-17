package jay

func ReadUintVariable(y []byte) (uint, int) {
	switch y[_0] {
	case _0:
		return _0, _1
	case _1:
		return uint(y[_1]), _2
	case _2:
		return uint(y[_1]) | uint(y[_2])<<_8, _3
	case _3:
		return uint(y[_1]) | uint(y[_2])<<_8 | uint(y[_3])<<_16, _4
	case _4:
		return uint(y[_1]) | uint(y[_2])<<_8 | uint(y[_3])<<_16 | uint(y[_4])<<_24, _5
	case _5:
		return uint(y[_1]) | uint(y[_2])<<_8 | uint(y[_3])<<_16 | uint(y[_4])<<_24 |
			uint(y[_5])<<_32, _6
	case _6:
		return uint(y[_1]) | uint(y[_2])<<_8 | uint(y[_3])<<_16 | uint(y[_4])<<_24 |
			uint(y[_5])<<_32 | uint(y[_6])<<_40, _7
	case _7:
		return uint(y[_1]) | uint(y[_2])<<_8 | uint(y[_3])<<_16 | uint(y[_4])<<_24 |
			uint(y[_5])<<_32 | uint(y[_6])<<_40 | uint(y[_7])<<_48, _8
	default:
		return uint(y[_1]) | uint(y[_2])<<_8 | uint(y[_3])<<_16 | uint(y[_4])<<_24 |
			uint(y[_5])<<_32 | uint(y[_6])<<_40 | uint(y[_7])<<_48 | uint(y[_8])<<_56, 9
	}
}

// WriteUintVariableAt ...
func WriteUintVariableAt(y []byte, u uint, length uint8, at int) int {
	WriteUintVariable(y, u, length)
	return at + int(length)
}

// WriteUintVariable ...
func WriteUintVariable(y []byte, u uint, length uint8) {
	y[_0] = length
	switch length {
	// case 1: No further processing required.
	case _2:
		y[_1] = byte(u)
	case _3:
		y[_1], y[_2] = byte(u), byte(u>>_8)
	case _4:
		y[_1], y[_2], y[_3] = byte(u), byte(u>>_8), byte(u>>_16)
	case _5:
		y[_1], y[_2], y[_3], y[_4] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
	case _6:
		y[_1], y[_2], y[_3], y[_4], y[_5] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32)
	case _7:
		y[_1], y[_2], y[_3], y[_4], y[_5], y[_6] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40)
	case _8:
		y[_1], y[_2], y[_3], y[_4], y[_5], y[_6], y[_7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48)
	case 9:
		y[_1], y[_2], y[_3], y[_4], y[_5], y[_6], y[_7], y[_8] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
	}
}

/*// Deprecated
// lenUintDeprecated ...
func lenUintDeprecated(u uint) uint8 {
	switch {
	case u == 0:
		return 1
	case u <= MaxUint8:
		return 2
	case u <= MaxUint16:
		return 3
	case u <= MaxUint24:
		return 4
	case u <= MaxUint32:
		return 5
	case u <= MaxUint40:
		return 6
	case u <= MaxUint48:
		return 7
	case u <= MaxUint56:
		return 8
	default:
		return 9
	}
}*/

// WriteUint24 ...
func WriteUint24(y []byte, u uint64) {
	y[_0], y[_1], y[_2] = byte(u), byte(u>>_8), byte(u>>_16)
}

// WriteUint40 ...
func WriteUint40(y []byte, u uint64) {
	y[_0], y[_1], y[_2], y[_3], y[_4] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32)
}

// WriteUint48 ...
func WriteUint48(y []byte, u uint64) {
	y[_0], y[_1], y[_2], y[_3], y[_4], y[_5] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40)
}

// WriteUint56 ...
func WriteUint56(y []byte, u uint64) {
	y[_0], y[_1], y[_2], y[_3], y[_4], y[_5], y[_6] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48)
}
