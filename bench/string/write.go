package string

func write0(y []byte, s string) {
	y[0] = byte(len(s))
	copy(y[1:], s)
}

func write0F(y []byte, s string) {
	y[0] = byte(len(s))
	if y[0] != 0 {
		copy(y[1:], s)
	}
}

func write1(y []byte, s string) {
	y[0] = byte(len(s))
	copy(y[1:y[0]+1], s)
}

func write1F(y []byte, s string) {
	y[0] = byte(len(s))
	if y[0] != 0 {
		copy(y[1:y[0]+1], s)
	}
}

func write2(y []byte, s string, length int) {
	y[0] = byte(length)
	copy(y[1:], s)
}

func write2F(y []byte, s string, length int) {
	y[0] = byte(length)
	if length != 0 {
		copy(y[1:], s)
	}
}

func write3(y []byte, s string, length int) {
	y[0] = byte(length)
	copy(y[1:length+1], s)
}

func write3F(y []byte, s string, length int) {
	y[0] = byte(length)
	if length != 0 {
		copy(y[1:length+1], s)
	}
}

func write2FC(y []byte, s string, length int) {
	y[0] = byte(length)
	if length != 0 {
		copy(y[strSizeOf:], s)
	}
}

func writeAt(y []byte, s string, length, at int) int {
	y[0] = byte(length) // Set how long the string is.
	if length != 0 {
		copy(y[strSizeOf:length+strSizeOf], s)
	}
	return at + length + strSizeOf
}

func writeAt2(y []byte, s string, length, at int) int {
	y[0] = byte(length) // Set how long the string is.
	if length != 0 {
		copy(y[strSizeOf:], s)
	}
	return at + length + strSizeOf
}

func WriteStringAtV2(y []byte, s string, length, at int) int {
	y[0] = byte(length) // Set how long the string is.
	if length != 0 {

	}
	return 0
}
