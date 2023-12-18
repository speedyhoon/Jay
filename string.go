package jay

const strSizeOf = 1

func WriteString(y []byte, s string, length int) {
	y[0] = byte(length) // Set how long the string is.
	if length != 0 {
		copy(y[strSizeOf:length+strSizeOf], s)
	}
}

func WriteStringAt(y []byte, s string, length, at int) int {
	y[0] = byte(length) // Set how long the string is.
	if length != 0 {
		copy(y[strSizeOf:length+strSizeOf], s)
	}
	return at + length + strSizeOf
}

func WriteStringAtPtr(y []byte, s *string, length, at int) int {
	y[0] = byte(length) // Set how long the string is.
	if length != 0 {
		copy(y[strSizeOf:length+strSizeOf], *s)
	}
	return at + length + strSizeOf
}

func ReadString(b []byte) (h string, size int, _ bool) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf || len(b) < size {
		return "", size, size == strSizeOf
	}

	return string(b[strSizeOf:size]), size, true
}

func ReadStringPtrErr(b []byte, h *string) error {
	size := int(b[0]) + strSizeOf
	if size == strSizeOf {
		return nil
	}
	if len(b) < size {
		return ErrUnexpectedEOB
	}

	*h = string(b[strSizeOf:size])
	return nil
}

func ReadStringAt(b []byte, at int) (h string, size int, ok bool) {
	h, size, ok = ReadString(b)
	return h, at + size, ok
}

func WriteStrings(y []byte, s []string) {
	y[0] = byte(len(s)) // Set slice qty.
	if y[0] == 0 {
		return
	}

	index := 1
	for i := range s {
		WriteString(y[index:], s[i], len(s[i]))
		index += len(s[i]) + strSizeOf
	}
}

func ReadStrings(b []byte) (h []string, size int, ok bool) {
	l := len(b)
	if l <= strSizeOf || b[0] == 0 {
		return
	}

	size = 1
	var z int
	h = make([]string, b[0])
	for i := uint8(0); i < b[0]; i++ {
		if size >= l {
			return
		}
		h[i], z, ok = ReadString(b[size:])
		size += z
		if !ok {
			return
		}
	}

	return

	//size = int(b[0]) + strSizeOf
	//if size == strSizeOf || len(b) < size {
	//	return
	//}
	//
	//return string(b[strSizeOf:size]), size, true
}

func ReadStringPtr(b []byte, h *string) (size int, _ bool) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf || len(b) < size {
		return
	}

	*h = string(b[strSizeOf:size])
	return size, true
}
