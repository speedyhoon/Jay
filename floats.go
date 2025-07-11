package jay

import "math"

func WriteFloat32s(y []byte, slice []float32, length int) {
	if length == 0 {
		return
	}

	for i, f := range slice {
		u := math.Float32bits(f)
		y[i*_4], y[i*_4+1], y[i*_4+2], y[i*_4+3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
	}
}

func WriteFloat64s(y []byte, slice []float64, length int) {
	if length == 0 {
		return
	}

	for i := range slice {
		WriteFloat64(y[i*_8:i*_8+_8], slice[i])
	}
}

func ReadFloat32s(y []byte, length int) (t []float32) {
	if length == 0 {
		return
	}

	t = make([]float32, length)
	for i, c := 0, 0; i < length; i, c = i+1, c+_4 {
		t[i] = math.Float32frombits(ReadUint32(y[c : c+4]))
	}
	return
}

func ReadFloat64s(y []byte, length int) (t []float64) {
	if length == 0 {
		return
	}

	t = make([]float64, length)
	for i := 0; i < length; i++ {
		t[i] = math.Float64frombits(ReadUint64(y[i*_8 : i*_8+_8]))
	}
	return
}
