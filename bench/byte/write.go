package byte

func write0(y []byte, s []byte) {
	y[0] = byte(len(s))
	copy(y[1:], s)
}

func write0F(y []byte, s []byte) {
	y[0] = byte(len(s))
	if y[0] != 0 {
		copy(y[1:], s)
	}
}

func write1(y []byte, s []byte) {
	y[0] = byte(len(s))
	copy(y[1:y[0]+1], s)
}

func write1F(y []byte, s []byte) {
	y[0] = byte(len(s))
	if y[0] != 0 {
		copy(y[1:y[0]+1], s)
	}
}

func write2(y []byte, s []byte, length int) {
	y[0] = byte(length)
	copy(y[1:length+1], s)
}

func write2F(y []byte, s []byte, length int) {
	y[0] = byte(length)
	if length != 0 {
		copy(y[1:length+1], s)
	}
}

func write2Fb(y []byte, s []byte, length int) {
	y[0] = byte(length)
	if y[0] != 0 {
		copy(y[1:length+1], s)
	}
}

func write3(y []byte, s []byte, length int) {
	y[0] = byte(length)
	copy(y[1:], s)
}

func write3F(y []byte, s []byte, length int) {
	y[0] = byte(length)
	if length != 0 {
		copy(y[1:], s)
	}
}

func write3Fb(y []byte, s []byte, length int) {
	y[0] = byte(length)
	if y[0] != 0 {
		copy(y[1:], s)
	}
}

func writeAt(y []byte, s []byte, length, at int) int {
	y[0] = byte(length)
	copy(y[1:], s)
	return length + at
}

func writeAtF(y []byte, s []byte, length, at int) int {
	y[0] = byte(length)
	if length != 0 {
		copy(y[1:], s)
	}
	return length + at
}

func writeAtFb(y []byte, s []byte, length, at int) int {
	y[0] = byte(length)
	if y[0] != 0 {
		copy(y[1:], s)
	}
	return length + at
}
