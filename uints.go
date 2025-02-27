package jay

func WriteUintsX32(y []byte, slice []uint) {
	for i := range slice {
		WriteUintX32(y[i*_4:i*_4+_4], slice[i])
	}
}

func ReadUintsX32(y []byte, length int) (t []uint) {
	if length == 0 {
		return
	}

	t = make([]uint, length)
	for i := 0; i < length; i++ {
		t[i] = ReadUintX32(y[i*_4 : i*_4+_4])
	}
	return
}

func WriteUintsX64(y []byte, slice []uint) {
	for i := range slice {
		WriteUintX64(y[i*_8:i*_8+_8], slice[i])
	}
}

func ReadUintsX64(y []byte, length int) (t []uint) {
	if length == 0 {
		return
	}

	t = make([]uint, length)
	for i := 0; i < length; i++ {
		t[i] = ReadUintX64(y[i*_8 : i*_8+_8])
	}
	return
}

func WriteUint64s(y []byte, slice []uint64) {
	for i := range slice {
		WriteUint64(y[i*_8:i*_8+_8], slice[i])
	}
}

func ReadUint64s(y []byte, length int) (t []uint64) {
	if length == 0 {
		return
	}

	t = make([]uint64, length)
	for i := 0; i < length; i++ {
		t[i] = ReadUint64(y[i*_8 : i*_8+_8])
	}
	return
}

func WriteUint32s(y []byte, slice []uint32) {
	for i := range slice {
		WriteUint32(y[i*_4:i*_4+_4], slice[i])
	}
}

func ReadUint32s(y []byte, length int) (t []uint32) {
	if length == 0 {
		return
	}

	t = make([]uint32, length)
	for i := 0; i < length; i++ {
		t[i] = ReadUint32(y[i*_4 : i*_4+_4])
	}
	return
}

func WriteUint16s(y []byte, ints []uint16, length int) {
	if length == 0 {
		return
	}

	for i := 0; i < length; i++ {
		y[i*2], y[i*2+1] = byte(ints[i]), byte(ints[i]>>_8)
	}
}

func ReadUint16s(y []byte, length int) (t []uint16) {
	if length == 0 {
		return
	}

	t = make([]uint16, length)
	for i := 0; i < length; i++ {
		t[i] = uint16(y[i*2]) | uint16(y[i*2+1])<<_8
	}
	return
}
