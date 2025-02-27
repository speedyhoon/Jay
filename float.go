package jay

import "math"

func WriteFloat32(y []byte, f float32) {
	WriteUint32(y, math.Float32bits(f))
}

func WriteFloat64(y []byte, f float64) {
	WriteUint64(y, math.Float64bits(f))
}

func ReadFloat32(y []byte) float32 {
	return math.Float32frombits(ReadUint32(y))
}

func ReadFloat64(y []byte) float64 {
	return math.Float64frombits(ReadUint64(y))
}
