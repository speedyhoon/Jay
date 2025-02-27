package jay

// ReadUint64 converts the first 8 bits in `b` into a uint64.
func ReadUint64(y []byte) uint64 {
	return uint64(y[0]) | uint64(y[1])<<8 | uint64(y[2])<<16 | uint64(y[3])<<24 |
		uint64(y[4])<<32 | uint64(y[5])<<40 | uint64(y[6])<<48 | uint64(y[7])<<56
}

// ReadUintX64 converts the first 8 bits in `b` into a uint.
func ReadUintX64(y []byte) uint {
	return uint(y[0]) | uint(y[1])<<8 | uint(y[2])<<16 | uint(y[3])<<24 |
		uint(y[4])<<32 | uint(y[5])<<40 | uint(y[6])<<48 | uint(y[7])<<56
}

// ReadUintX32 converts the first 6 bits in `b` into a uint.
func ReadUintX32(y []byte) uint {
	return uint(y[0]) | uint(y[1])<<8 | uint(y[2])<<16 | uint(y[3])<<24
}

// ReadUint32 converts the first 4 bits in `b` into a uint32.
func ReadUint32(y []byte) uint32 {
	return uint32(y[0]) | uint32(y[1])<<8 | uint32(y[2])<<16 | uint32(y[3])<<24
}

// ReadUint16 converts the first 2 bits in `b` into a uint16.
func ReadUint16(y []byte) uint16 {
	return uint16(y[0]) | uint16(y[1])<<8
}

// WriteUint64 sets the first 8 bits in `b` to the value of `i`.
func WriteUint64(y []byte, u uint64) {
	y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
}

// WriteUintX64 sets the first 8 bits in `b` to the value of `i`.
func WriteUintX64(y []byte, u uint) {
	y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
}

// WriteUintX32 sets the first 4 bits in `b` to the value of `i`.
func WriteUintX32(y []byte, u uint) {
	y[0], y[1], y[2], y[3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
}

// WriteUint32 sets the first 4 bits in `b` to the value of `i`.
func WriteUint32(y []byte, u uint32) {
	y[0], y[1], y[2], y[3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
}

// WriteUint16 sets the first 2 bits in `b` to the value of `i`.
func WriteUint16(y []byte, u uint16) {
	y[0], y[1] = byte(u), byte(u>>_8)
}
