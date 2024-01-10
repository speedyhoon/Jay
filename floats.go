package jay

import "math"

func WriteFloat32s(b []byte, slice []float32, length int) {
	if length == 0 {
		return
	}

	b[0] = byte(length)
	for i, f := range slice {
		u := math.Float32bits(f)

		b[i*_4+1], b[i*_4+2], b[i*_4+3], b[i*_4+4] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
	}
}

func WriteFloat64s(b []byte, slice []float64, length int) {
	if length == 0 {
		return
	}

	b[0] = byte(length)
	for i, f := range slice {
		WriteFloat64(b[i*_8+_1:i*_8+9], f)
	}
}

func ReadFloat32s(b []byte) (t []float32) {
	if b[0] == 0 {
		return
	}

	t = make([]float32, b[0])
	for i, c, d := byte(0), byte(1), byte(5); i < b[0]; c, d, i = c+_4, d+_4, i+1 {
		t[i] = math.Float32frombits(ReadUint32(b[c:d]))
	}
	return
}

func ReadFloat64s(b []byte) (t []float64) {
	if b[0] == 0 {
		return
	}

	t = make([]float64, b[0])
	for i := byte(0); i < b[0]; i++ {
		t[i] = math.Float64frombits(ReadUint64(b[i*_8+_1 : i*_8+9]))
	}
	return
}
