package jay

import "math"

func WriteFloat32(b []byte, f float32) {
	WriteUint32(b, math.Float32bits(f))
}

func WriteFloat64(b []byte, f float64) {
	WriteUint64(b, math.Float64bits(f))
}
