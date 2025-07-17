package jay

func WriteIntsX32(y []byte, slice []int) {
	for i := range slice {
		WriteIntX32(y[i*_4:i*_4+_4], slice[i])
	}
}

func ReadIntsX32(y []byte, length int) (t []int) {
	if length == _0 {
		return
	}

	t = make([]int, length)
	for i := _0; i < length; i++ {
		t[i] = ReadIntX32(y[i*_4 : i*_4+_4])
	}
	return
}

func WriteIntsX64(y []byte, slice []int) {
	for i := range slice {
		WriteIntX64(y[i*_8:i*_8+_8], slice[i])
	}
}

func ReadIntsX64(y []byte, length int) (t []int) {
	if length == _0 {
		return
	}

	t = make([]int, length)
	for i := _0; i < length; i++ {
		t[i] = ReadIntX64(y[i*_8 : i*_8+_8])
	}
	return
}

func WriteInt64s(y []byte, slice []int64) {
	for i := range slice {
		WriteInt64(y[i*_8:i*_8+_8], slice[i])
	}
}

func ReadInt64s(y []byte, length int) (t []int64) {
	if length == _0 {
		return
	}

	t = make([]int64, length)
	for i := _0; i < length; i++ {
		t[i] = ReadInt64(y[i*_8 : i*_8+_8])
	}
	return
}

func WriteInt32s(y []byte, slice []int32) {
	for i := range slice {
		WriteInt32(y[i*_4:i*_4+_4], slice[i])
	}
}

func ReadInt32s(y []byte, length int) (t []int32) {
	if length == _0 {
		return
	}

	t = make([]int32, length)
	for i := _0; i < length; i++ {
		t[i] = ReadInt32(y[i*_4 : i*_4+_4])
	}
	return
}

func WriteInt8s(y []byte, slice []int8) {
	for i, v := range slice {
		y[i] = byte(v)
	}
}

func ReadInt8s(y []byte, length int) (t []int8) {
	if length == _0 {
		return
	}

	t = make([]int8, length)
	for i := _0; i < length; i++ {
		t[i] = int8(y[i])
	}
	return
}

func WriteInt16s(y []byte, slice []int16, length int) {
	if length == _0 {
		return
	}

	for i := _0; i < length; i++ {
		y[i*_2], y[i*_2+_1] = byte(slice[i]), byte(slice[i]>>_8)
	}
}

func ReadInt16s(y []byte, length int) (t []int16) {
	if length == _0 {
		return
	}

	t = make([]int16, length)
	for i := _0; i < length; i++ {
		t[i] = int16(y[i*_2]) | int16(y[i*_2+_1])<<_8
	}
	return
}
