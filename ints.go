package jay

func WriteInt8s(b []byte, slice []int8) {
	for i, v := range slice {
		b[i] = byte(v)
	}
}

func ReadInt8s(b []byte, length int) (t []int8) {
	if length == 0 {
		return
	}

	t = make([]int8, length)
	for i := 0; i < length; i++ {
		t[i] = int8(b[i])
	}
	return
}

func WriteInt16s(b []byte, ints []int16, length int) {
	if length == 0 {
		return
	}

	for i := 0; i < length; i++ {
		b[i*2], b[i*2+1] = byte(ints[i]), byte(ints[i]>>_8)
	}
}

func ReadInt16s(b []byte, length int) (t []int16) {
	if length == 0 {
		return
	}

	t = make([]int16, length)
	for i := 0; i < length; i++ {
		t[i] = int16(b[i*2]) | int16(b[i*2+1])<<_8
	}
	return
}
