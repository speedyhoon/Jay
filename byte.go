package jay

func WriteBytes(b []byte, y []byte) {
	b[0] = byte(len(y))
	if b[0] >= 1 {
		copy(b[1:], y)
	}
}

func ReadBytes(b []byte) ([]byte, bool) {
	if len(b) < int(b[0]+1) {
		return nil, false
	}

	return b[1 : b[0]+1], true
}
