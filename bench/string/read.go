package string

const strSizeOf = 1

func ReadStringW(b []byte) (h string, size int, _ bool) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf {
		return "", size, true
	}

	if len(b) < size {
		return "", size, false
	}

	return string(b[strSizeOf:size]), size, true
}

func read5(b []byte) (h string, size int) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf || len(b) < size {
		return "", -1
	}

	return string(b[strSizeOf:size]), size
}

func ReadStringAtW(b []byte, at int) (h string, size int, ok bool) {
	h, size, ok = ReadStringW(b)
	return h, at + size, ok
}

func ReadStringAtV5(b []byte, at int) (h string, size int) {
	size = int(b[0]) + strSizeOf
	if size == strSizeOf || len(b) < size {
		return "", -1
	}

	return string(b[strSizeOf:size]), at + size
}
