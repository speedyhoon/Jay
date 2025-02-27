package jay

// StringsSize8 calculates the total size needed to represent a []string and
// returns the quantity of strings in a slice of strings
// StringsSize8 only supports up to 255 items.
func StringsSize8(s []string) (total, qty int, sizes []byte) {
	qty = len(s)
	sizes = make([]byte, qty)
	var l int
	for i := range s {
		l = len(s[i])
		sizes[i] = byte(l)
		total += l
	}
	return
}

// StringsSize16 only supports up to 65,535 items.
func StringsSize16(s []string) (total, qty int, sizes []byte) {
	qty = len(s)
	sizes = make([]byte, qty*2)
	var l int
	for i := range s {
		l = len(s[i])
		sizes[i*2], sizes[i*2+1] = byte(l), byte(l>>8)
		total += l
	}
	return
}

// StringsSize24 only supports up to 16,777,215 items.
func StringsSize24(s []string) (total, qty int, sizes []byte) {
	qty = len(s)
	sizes = make([]byte, qty*3)
	var l int
	for i := range s {
		l = len(s[i])
		sizes[i*3], sizes[i*3+1], sizes[i*3+2] = byte(l), byte(l>>8), byte(l>>16)
		total += l
	}
	return
}

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

/*func ReadStrings(y []byte) (s []string, size int, ok bool) {
	l := len(y)
	if l <= strSizeOf || y[0] == 0 {
		return
	}

	size = 1
	var z int
	s = make([]string, y[0])
	for i := uint8(0); i < y[0]; i++ {
		if size >= l {
			return
		}
		s[i], z, ok = ReadString(y[size:])
		size += z
		if !ok {
			return
		}
	}

	return

	//size = int(y[0]) + strSizeOf
	//if size == strSizeOf || len(y) < size {
	//	return
	//}
	//
	//return string(y[strSizeOf:size]), size, true
}

func ReadStringPtr(y []byte, h *string) (size int, ok bool) {
	size = int(y[0]) + strSizeOf
	if size == strSizeOf || len(y) < size {
		return
	}

	*h = string(y[strSizeOf:size])
	return size, true
}*/
