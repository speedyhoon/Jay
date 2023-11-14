package jay

const (
	MaxUint8  = 1<<8 - 1  // 255
	MaxUint16 = 1<<16 - 1 // 65535
	MaxUint24 = 1<<24 - 1 // 16777215
	MaxUint32 = 1<<32 - 1 // 4294967295
	MaxUint40 = 1<<40 - 1 // 1099511627775
	MaxUint48 = 1<<48 - 1 // 281474976710655
	MaxUint56 = 1<<56 - 1 // 72057594037927935
	//MaxUint64 = 1<<64 - 1 // 18446744073709551615
)

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

// WriteUint ...
func WriteUint(b []byte, u uint, length uint8) {
	b[0] = length
	switch length {
	// case 1: No further processing required.
	case 2:
		b[1] = byte(u)
	case 3:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
	case 4:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
	case 5:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
	case 6:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		b[5] = byte(u >> 32)
	case 7:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		b[5] = byte(u >> 32)
		b[6] = byte(u >> 40)
	case 8:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		b[5] = byte(u >> 32)
		b[6] = byte(u >> 40)
		b[7] = byte(u >> 48)
	case 9:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		b[5] = byte(u >> 32)
		b[6] = byte(u >> 40)
		b[7] = byte(u >> 48)
		b[8] = byte(u >> 56)
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
		b[2] = byte(u >> 8)
		return 3
	case 3:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		return 4
	case 4:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		return 5
	case 5:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		b[5] = byte(u >> 32)
		return 6
	case 6:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		b[5] = byte(u >> 32)
		b[6] = byte(u >> 40)
		return 7
	case 7:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		b[5] = byte(u >> 32)
		b[6] = byte(u >> 40)
		b[7] = byte(u >> 48)
		return 8
	case 8:
		b[1] = byte(u)
		b[2] = byte(u >> 8)
		b[3] = byte(u >> 16)
		b[4] = byte(u >> 24)
		b[5] = byte(u >> 32)
		b[6] = byte(u >> 40)
		b[7] = byte(u >> 48)
		b[8] = byte(u >> 56)
		return 9
	}
}

// WriteUint64 ...
func WriteUint64(b []byte, u uint64) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(u), byte(u>>8), byte(u>>16), byte(u>>24), byte(u>>32), byte(u>>40), byte(u>>48), byte(u>>56)
}

// WriteUint56 ...
func WriteUint56(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
	b[3] = byte(u >> 24)
	b[4] = byte(u >> 32)
	b[5] = byte(u >> 40)
	b[6] = byte(u >> 48)
}

// WriteUint48 ...
func WriteUint48(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
	b[3] = byte(u >> 24)
	b[4] = byte(u >> 32)
	b[5] = byte(u >> 40)
}

// WriteUint40 ...
func WriteUint40(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
	b[3] = byte(u >> 24)
	b[4] = byte(u >> 32)
}

// WriteUint32 ...
func WriteUint32(b []byte, u uint32) {
	b[0], b[1], b[2], b[3] = byte(u), byte(u>>8), byte(u>>16), byte(u>>24)
}

// WriteUint24 ...
func WriteUint24(b []byte, u uint64) {
	b[0] = byte(u)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
}

// WriteUint16 ...
func WriteUint16(b []byte, u uint16) {
	b[0], b[1] = byte(u), byte(u>>8)
}
