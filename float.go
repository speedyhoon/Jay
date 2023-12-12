package jay

import "math"

func WriteFloat32(b []byte, f float32) {
	WriteUint32(b, math.Float32bits(f))
}

func WriteFloat64(b []byte, f float64) {
	WriteUint64(b, math.Float64bits(f))
}

func ReadFloat32(b []byte) float32 {
	return math.Float32frombits(ReadUint32(b))
}

func ReadFloat64(b []byte) float64 {
	return math.Float64frombits(ReadUint64(b))
}
