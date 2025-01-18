package jay

//const strSizeOf = 1

func WriteStrings(ln, y []byte, s []string) {
	var l, qty int
	for i := range s {
		l = len(s[i])
		ln[i] = byte(l)
		copy(y[qty:qty+l], s[i])
		qty += l
	}
}

/*func ReadStrings(b []byte) (h []string, size int, ok bool) {
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
}*/
