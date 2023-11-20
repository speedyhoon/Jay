package jay

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

// DEPRECATED
func WriteUintDEPRECATED(b []byte, u uint) int {
	b[0] = LenUint(u) - 1
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

// LenUint ...
func LenUint(u uint) uint8 {
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
}

// WriteUint24 ...
func WriteUint24(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> _8)
	b[2] = byte(u >> _16)
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

// WriteUint48 ...
func WriteUint48(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> _8)
	b[2] = byte(u >> _16)
	b[3] = byte(u >> _24)
	b[4] = byte(u >> _32)
	b[5] = byte(u >> _40)
}

// WriteUint40 ...
func WriteUint40(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> _8)
	b[2] = byte(u >> _16)
	b[3] = byte(u >> _24)
	b[4] = byte(u >> _32)
}
