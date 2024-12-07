package jay

func WriteBytes(y []byte, s []byte, length int) {
	if length != 0 {
		y[0] = byte(length) // Set how long the []byte is.
		copy(y[strSizeOf:length+strSizeOf], s)
	}
}

func ReadBytes(b []byte) (h []byte, size int, _ bool) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf || len(b) < size {
		return nil, size, size == strSizeOf
	}

	return b[strSizeOf:size], size, true
}
