package jay

// LenStr returns the length of the string pointer.
func LenStr(s *string) int {
	if s == nil {
		return 0
	}

	return len(*s)
}

// Constants `on` [11000000] & `off` [10000000] are these values for easier bit manipulation & shifting (see `bitIndex`).
const on, off uint8 = 192, 128

// WriteBoolPtr converts several bool pointers to []byte.
// TODO add support to jay/generate for *bool.
func WriteBoolPtr(x ...*bool) (b []byte) {
	// Assume all *bools need 2 bits.
	b = make([]byte, (len(x)+3)/4)

	var byteIndex, bitIndex uint8

	for _, y := range x {
		if bitIndex >= 8 {
			bitIndex %= 8
			byteIndex++
		}

		switch {
		case y == nil:
			bitIndex++
		case *y:
			b[byteIndex] |= on >> bitIndex
			if bitIndex == 7 {
				// Right-shifting `on` 7 bits splits the bits between this byte at `byteIndex` and the next byte, e.g.: [00000001,10000000].
				// Using an or operation with `off` is simpler than calculating with bit shift operators.
				b[byteIndex+1] |= off
			}
			bitIndex += 2
		default:
			b[byteIndex] |= off >> bitIndex
			bitIndex += 2
		}
	}

	return b[:byteIndex+1]
}

// ReadBoolPtr converts []byte to a list of bool pointers.
// TODO add support to jay/generate for *bool.
func ReadBoolPtr(b []byte) (x []*bool) {
	l := uint8(len(b))
	x = make([]*bool, l*8)

	var byteIndex, bitIndex uint8
	for c := 0; byteIndex < l; {
		if isBitTrue(b[byteIndex], &bitIndex) {
			val := isBitTrue(b[byteIndex], &bitIndex)
			x[c] = &val
		}
		c++

		if bitIndex >= 8 {
			bitIndex %= 8
			byteIndex++
		}
	}

	return
}

// isBitTrue checks if a single bit at index bitIndex isn't zero and increments `bitIndex`.
func isBitTrue(b byte, bitIndex *uint8) bool {
	compareTo := off >> *bitIndex
	*bitIndex++
	return b&compareTo == compareTo
}
