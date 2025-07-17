package jay

// ReadUint64 converts the first 8 bits in `b` into a uint64.
func ReadUint64(y []byte) uint64 {
	return uint64(y[_0]) | uint64(y[_1])<<_8 | uint64(y[_2])<<_16 | uint64(y[_3])<<_24 |
		uint64(y[_4])<<_32 | uint64(y[_5])<<_40 | uint64(y[_6])<<_48 | uint64(y[_7])<<_56
}

// ReadUintX64 converts the first 8 bits in `b` into a uint.
func ReadUintX64(y []byte) uint {
	return uint(y[_0]) | uint(y[_1])<<_8 | uint(y[_2])<<_16 | uint(y[_3])<<_24 |
		uint(y[_4])<<_32 | uint(y[_5])<<_40 | uint(y[_6])<<_48 | uint(y[_7])<<_56
}

// ReadUintX32 converts the first 6 bits in `b` into a uint.
func ReadUintX32(y []byte) uint {
	return uint(y[_0]) | uint(y[_1])<<_8 | uint(y[_2])<<_16 | uint(y[_3])<<_24
}

// ReadUint32 converts the first 4 bits in `b` into a uint32.
func ReadUint32(y []byte) uint32 {
	return uint32(y[_0]) | uint32(y[_1])<<_8 | uint32(y[_2])<<_16 | uint32(y[_3])<<_24
}

// ReadUint16 converts the first 2 bits in `b` into a uint16.
func ReadUint16(y []byte) uint16 {
	return uint16(y[_0]) | uint16(y[_1])<<_8
}

// WriteUint64 sets the first 8 bits in `b` to the value of `i`.
func WriteUint64(y []byte, u uint64) {
	y[_0], y[_1], y[_2], y[_3], y[_4], y[_5], y[_6], y[_7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
}

// WriteUintX64 sets the first 8 bits in `b` to the value of `i`.
func WriteUintX64(y []byte, u uint) {
	y[_0], y[_1], y[_2], y[_3], y[_4], y[_5], y[_6], y[_7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
}

// WriteUintX32 sets the first 4 bits in `b` to the value of `i`.
func WriteUintX32(y []byte, u uint) {
	y[_0], y[_1], y[_2], y[_3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
}

// WriteUint32 sets the first 4 bits in `b` to the value of `i`.
func WriteUint32(y []byte, u uint32) {
	y[_0], y[_1], y[_2], y[_3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
}

// WriteUint16 sets the first 2 bits in `b` to the value of `i`.
func WriteUint16(y []byte, u uint16) {
	y[_0], y[_1] = byte(u), byte(u>>_8)
}
