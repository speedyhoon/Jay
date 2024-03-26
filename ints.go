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
