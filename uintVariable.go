package jay

func ReadUintVariable(y []byte) (uint, int) {
	switch y[0] {
	case 0:
		return 0, 1
	case 1:
		return uint(y[1]), 2
	case 2:
		return uint(y[1]) | uint(y[2])<<_8, 3
	case 3:
		return uint(y[1]) | uint(y[2])<<_8 | uint(y[3])<<_16, 4
	case 4:
		return uint(y[1]) | uint(y[2])<<_8 | uint(y[3])<<_16 | uint(y[4])<<_24, 5
	case 5:
		return uint(y[1]) | uint(y[2])<<_8 | uint(y[3])<<_16 | uint(y[4])<<_24 |
			uint(y[5])<<_32, 6
	case 6:
		return uint(y[1]) | uint(y[2])<<_8 | uint(y[3])<<_16 | uint(y[4])<<_24 |
			uint(y[5])<<_32 | uint(y[6])<<_40, 7
	case 7:
		return uint(y[1]) | uint(y[2])<<_8 | uint(y[3])<<_16 | uint(y[4])<<_24 |
			uint(y[5])<<_32 | uint(y[6])<<_40 | uint(y[7])<<_48, 8
	default:
		return uint(y[1]) | uint(y[2])<<_8 | uint(y[3])<<_16 | uint(y[4])<<_24 |
			uint(y[5])<<_32 | uint(y[6])<<_40 | uint(y[7])<<_48 | uint(y[8])<<_56, 9
	}
}

// WriteUintVariableAt ...
func WriteUintVariableAt(y []byte, u uint, length uint8, at int) int {
	WriteUintVariable(y, u, length)
	return at + int(length)
}

// WriteUintVariable ...
func WriteUintVariable(y []byte, u uint, length uint8) {
	y[0] = length
	switch length {
	// case 1: No further processing required.
	case 2:
		y[1] = byte(u)
	case 3:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
	case 4:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
	case 5:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
	case 6:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		y[5] = byte(u >> _32)
	case 7:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		y[5] = byte(u >> _32)
		y[6] = byte(u >> _40)
	case 8:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		y[5] = byte(u >> _32)
		y[6] = byte(u >> _40)
		y[7] = byte(u >> _48)
	case 9:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		y[5] = byte(u >> _32)
		y[6] = byte(u >> _40)
		y[7] = byte(u >> _48)
		y[8] = byte(u >> _56)
	}
}

/*// Deprecated
// WriteUintDeprecated ...
func WriteUintDeprecated(y []byte, u uint) int {
	y[0] = lenUintDeprecated(u) - 1
	switch y[0] {
	default:
		// Not required, value is zero.
		return 1
	case 1:
		y[1] = byte(u)
		return 2
	case 2:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		return 3
	case 3:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		return 4
	case 4:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		return 5
	case 5:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		y[5] = byte(u >> _32)
		return 6
	case 6:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		y[5] = byte(u >> _32)
		y[6] = byte(u >> _40)
		return 7
	case 7:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		y[5] = byte(u >> _32)
		y[6] = byte(u >> _40)
		y[7] = byte(u >> _48)
		return 8
	case 8:
		y[1] = byte(u)
		y[2] = byte(u >> _8)
		y[3] = byte(u >> _16)
		y[4] = byte(u >> _24)
		y[5] = byte(u >> _32)
		y[6] = byte(u >> _40)
		y[7] = byte(u >> _48)
		y[8] = byte(u >> _56)
		return 9
	}
}

// Deprecated
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
	y[0], y[1], y[2] = byte(u), byte(u>>_8), byte(u>>_16)
}

// WriteUint40 ...
func WriteUint40(y []byte, u uint64) {
	y[0], y[1], y[2], y[3], y[4] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32)
}

// WriteUint48 ...
func WriteUint48(y []byte, u uint64) {
	y[0], y[1], y[2], y[3], y[4], y[5] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40)
}

// WriteUint56 ...
func WriteUint56(y []byte, u uint64) {
	y[0], y[1], y[2], y[3], y[4], y[5], y[6] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48)
}
