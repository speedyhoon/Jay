package jay

import (
	"errors"
)

// ErrUnexpectedEOB indicates the end of the byte slice was unexpectedly encountered
// while deserializing a fixed-size block, resulting in an incomplete result.
var ErrUnexpectedEOB = errors.New("unexpected EOB")

func ByteToBool(b byte) bool {
	return b == 1
}

func ReadUint64(buf []byte) uint64 {
	return uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 |
		uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56
}

func ReadUint(b []byte) (uint, int) {
	//bits.Len()
	switch b[0] {
	case 0:
		return 0, 1
	case 1:
		return uint(b[1]), 2
	case 2:
		return uint(b[1]) | uint(b[2])<<8, 3
	case 3:
		return uint(b[1]) | uint(b[2])<<8 | uint(b[3])<<16, 4
	case 4:
		return uint(b[1]) | uint(b[2])<<8 | uint(b[3])<<16 | uint(b[4])<<24, 5
	case 5:
		return uint(b[1]) | uint(b[2])<<8 | uint(b[3])<<16 | uint(b[4])<<24 |
			uint(b[5])<<32, 6
	case 6:
		return uint(b[1]) | uint(b[2])<<8 | uint(b[3])<<16 | uint(b[4])<<24 |
			uint(b[5])<<32 | uint(b[6])<<40, 7
	case 7:
		return uint(b[1]) | uint(b[2])<<8 | uint(b[3])<<16 | uint(b[4])<<24 |
			uint(b[5])<<32 | uint(b[6])<<40 | uint(b[7])<<48, 8
	default:
		return uint(b[1]) | uint(b[2])<<8 | uint(b[3])<<16 | uint(b[4])<<24 |
			uint(b[5])<<32 | uint(b[6])<<40 | uint(b[7])<<48 | uint(b[8])<<56, 9
	}
}
