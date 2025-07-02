package jay

func WriteComplex64(y []byte, c complex64) {
	WriteFloat32(y[:_4], real(c))
	WriteFloat32(y[_4:_8], imag(c))
}

func ReadComplex64(y []byte) complex64 {
	return complex(ReadFloat32(y[:_4]), ReadFloat32(y[_4:_8]))
}

func WriteComplex128(y []byte, c complex128) {
	WriteFloat64(y[:_8], real(c))
	WriteFloat64(y[_8:_16], imag(c))
}

func ReadComplex128(y []byte) complex128 {
	return complex(ReadFloat64(y[:_8]), ReadFloat64(y[_8:_16]))
}

func WriteComplex64s(y []byte, slice []complex64) {
	for i := range slice {
		WriteComplex64(y[i*_8:i*_8+_8], slice[i])
	}
}

func ReadComplex64s(y []byte, length int) (t []complex64) {
	if length == 0 {
		return
	}

	t = make([]complex64, length)
	for i := 0; i < length; i++ {
		t[i] = ReadComplex64(y[i*_8 : i*_8+_8])
	}
	return
}

func WriteComplex128s(y []byte, slice []complex128) {
	for i := range slice {
		WriteComplex128(y[i*_16:i*_16+_16], slice[i])
	}
}

func ReadComplex128s(y []byte, length int) (t []complex128) {
	if length == 0 {
		return
	}

	t = make([]complex128, length)
	for i := 0; i < length; i++ {
		t[i] = ReadComplex128(y[i*_16 : i*_16+_16])
	}
	return
}
