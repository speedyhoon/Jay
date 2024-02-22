package jay

func WriteBytes(y []byte, s []byte, length int) {
	if length != 0 {
		y[0] = byte(length) // Set how long the []byte is.
		copy(y[strSizeOf:length+strSizeOf], s)
	}
}

func WriteBytesAt(y []byte, s []byte, length, at int) int {
	if length != 0 {
		y[0] = byte(length) // Set how long the []byte is.
		copy(y[strSizeOf:], s)
	}
	return at + length + strSizeOf
}

func ReadBytes(b []byte) (h []byte, size int, _ bool) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf || len(b) < size {
		return nil, size, size == strSizeOf
	}

	return b[strSizeOf:size], size, true
}

func ReadBytesPtrErr(b []byte, h *[]byte) error {
	size := int(b[0]) + strSizeOf
	if size == strSizeOf {
		return nil
	}
	if len(b) < size {
		return ErrUnexpectedEOB
	}

	*h = b[strSizeOf:size]
	return nil
}

func ReadBytesAt(b []byte, at int) (h []byte, size int, ok bool) {
	h, size, ok = ReadBytes(b)
	return h, at + size, ok
}
