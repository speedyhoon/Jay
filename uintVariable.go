package jay

func ReadUintVariable(b []byte) (uint, int) {
	switch b[0] {
	case 0:
		return 0, 1
	case 1:
		return uint(b[1]), 2
	case 2:
		return uint(b[1]) | uint(b[2])<<_8, 3
	case 3:
		return uint(b[1]) | uint(b[2])<<_8 | uint(b[3])<<_16, 4
	case 4:
		return uint(b[1]) | uint(b[2])<<_8 | uint(b[3])<<_16 | uint(b[4])<<_24, 5
	case 5:
		return uint(b[1]) | uint(b[2])<<_8 | uint(b[3])<<_16 | uint(b[4])<<_24 |
			uint(b[5])<<_32, 6
	case 6:
		return uint(b[1]) | uint(b[2])<<_8 | uint(b[3])<<_16 | uint(b[4])<<_24 |
			uint(b[5])<<_32 | uint(b[6])<<_40, 7
	case 7:
		return uint(b[1]) | uint(b[2])<<_8 | uint(b[3])<<_16 | uint(b[4])<<_24 |
			uint(b[5])<<_32 | uint(b[6])<<_40 | uint(b[7])<<_48, 8
	default:
		return uint(b[1]) | uint(b[2])<<_8 | uint(b[3])<<_16 | uint(b[4])<<_24 |
			uint(b[5])<<_32 | uint(b[6])<<_40 | uint(b[7])<<_48 | uint(b[8])<<_56, 9
	}
}

// WriteUintVariable ...
func WriteUintVariable(b []byte, u uint, length uint8) {
	b[0] = length
	switch length {
	// case 1: No further processing required.
	case 2:
		b[1] = byte(u)
	case 3:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
	case 4:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
	case 5:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
	case 6:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		b[5] = byte(u >> _32)
	case 7:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		b[5] = byte(u >> _32)
		b[6] = byte(u >> _40)
	case 8:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		b[5] = byte(u >> _32)
		b[6] = byte(u >> _40)
		b[7] = byte(u >> _48)
	case 9:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		b[5] = byte(u >> _32)
		b[6] = byte(u >> _40)
		b[7] = byte(u >> _48)
		b[8] = byte(u >> _56)
	}
}

/*// Deprecated
// WriteUintDeprecated ...
func WriteUintDeprecated(b []byte, u uint) int {
	b[0] = lenUintDeprecated(u) - 1
	switch b[0] {
	default:
		// Not required, value is zero.
		return 1
	case 1:
		b[1] = byte(u)
		return 2
	case 2:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		return 3
	case 3:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		return 4
	case 4:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		return 5
	case 5:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		b[5] = byte(u >> _32)
		return 6
	case 6:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		b[5] = byte(u >> _32)
		b[6] = byte(u >> _40)
		return 7
	case 7:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		b[5] = byte(u >> _32)
		b[6] = byte(u >> _40)
		b[7] = byte(u >> _48)
		return 8
	case 8:
		b[1] = byte(u)
		b[2] = byte(u >> _8)
		b[3] = byte(u >> _16)
		b[4] = byte(u >> _24)
		b[5] = byte(u >> _32)
		b[6] = byte(u >> _40)
		b[7] = byte(u >> _48)
		b[8] = byte(u >> _56)
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
func WriteUint24(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> _8)
	b[2] = byte(u >> _16)
}

// WriteUint40 ...
func WriteUint40(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> _8)
	b[2] = byte(u >> _16)
	b[3] = byte(u >> _24)
	b[4] = byte(u >> _32)
}

// WriteUint48 ...
func WriteUint48(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> _8)
	b[2] = byte(u >> _16)
	b[3] = byte(u >> _24)
	b[4] = byte(u >> _32)
	b[5] = byte(u >> _40)
}

// WriteUint56 ...
func WriteUint56(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> _8)
	b[2] = byte(u >> _16)
	b[3] = byte(u >> _24)
	b[4] = byte(u >> _32)
	b[5] = byte(u >> _40)
	b[6] = byte(u >> _48)
}
