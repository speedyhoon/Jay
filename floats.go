package jay

import "math"

func WriteFloat32s(y []byte, slice []float32, length int) {
	if length == _0 {
		return
	}

	for i, f := range slice {
		u := math.Float32bits(f)
		y[i*_4], y[i*_4+_1], y[i*_4+_2], y[i*_4+_3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
	}
}

func WriteFloat64s(y []byte, slice []float64, length int) {
	if length == _0 {
		return
	}

	for i := range slice {
		WriteFloat64(y[i*_8:i*_8+_8], slice[i])
	}
}

func ReadFloat32s(y []byte, length int) (t []float32) {
	if length == _0 {
		return
	}

	t = make([]float32, length)
	for i, c := _0, _0; i < length; i, c = i+_1, c+_4 {
		t[i] = math.Float32frombits(ReadUint32(y[c : c+_4]))
	}
	return
}

func ReadFloat64s(y []byte, length int) (t []float64) {
	if length == _0 {
		return
	}

	t = make([]float64, length)
	for i := _0; i < length; i++ {
		t[i] = math.Float64frombits(ReadUint64(y[i*_8 : i*_8+_8]))
	}
	return
}
