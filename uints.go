package jay

func WriteUintsX32(b []byte, slice []uint) {
	for i := range slice {
		WriteUintX32(b[i*_4:i*_4+_4], slice[i])
	}
}

func ReadUintsX32(b []byte, length int) (t []uint) {
	if length == 0 {
		return
	}

	t = make([]uint, length)
	for i := 0; i < length; i++ {
		t[i] = ReadUintX32(b[i*_4 : i*_4+_4])
	}
	return
}

func WriteUintsX64(b []byte, slice []uint) {
	for i := range slice {
		WriteUintX64(b[i*_8:i*_8+_8], slice[i])
	}
}

func ReadUintsX64(b []byte, length int) (t []uint) {
	if length == 0 {
		return
	}

	t = make([]uint, length)
	for i := 0; i < length; i++ {
		t[i] = ReadUintX64(b[i*_8 : i*_8+_8])
	}
	return
}

func WriteUint64s(b []byte, slice []uint64) {
	for i := range slice {
		WriteUint64(b[i*_8:i*_8+_8], slice[i])
	}
}

func ReadUint64s(b []byte, length int) (t []uint64) {
	if length == 0 {
		return
	}

	t = make([]uint64, length)
	for i := 0; i < length; i++ {
		t[i] = ReadUint64(b[i*_8 : i*_8+_8])
	}
	return
}

func WriteUint32s(b []byte, slice []uint32) {
	for i := range slice {
		WriteUint32(b[i*_4:i*_4+_4], slice[i])
	}
}

func ReadUint32s(b []byte, length int) (t []uint32) {
	if length == 0 {
		return
	}

	t = make([]uint32, length)
	for i := 0; i < length; i++ {
		t[i] = ReadUint32(b[i*_4 : i*_4+_4])
	}
	return
}

func WriteUint16s(b []byte, ints []uint16, length int) {
	if length == 0 {
		return
	}

	for i := 0; i < length; i++ {
		b[i*2], b[i*2+1] = byte(ints[i]), byte(ints[i]>>_8)
	}
}

func ReadUint16s(b []byte, length int) (t []uint16) {
	if length == 0 {
		return
	}

	t = make([]uint16, length)
	for i := 0; i < length; i++ {
		t[i] = uint16(b[i*2]) | uint16(b[i*2+1])<<_8
	}
	return
}
