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
