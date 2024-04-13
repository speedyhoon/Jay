package jay

// ReadUint64 converts the first 8 bits in `b` into a uint64.
func ReadUint64(buf []byte) uint64 {
	return uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 |
		uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56
}

// ReadUintX64 converts the first 8 bits in `b` into a uint.
func ReadUintX64(buf []byte) uint {
	return uint(buf[0]) | uint(buf[1])<<8 | uint(buf[2])<<16 | uint(buf[3])<<24 |
		uint(buf[4])<<32 | uint(buf[5])<<40 | uint(buf[6])<<48 | uint(buf[7])<<56
}

// ReadUintX32 converts the first 6 bits in `b` into a uint.
func ReadUintX32(buf []byte) uint {
	return uint(buf[0]) | uint(buf[1])<<8 | uint(buf[2])<<16 | uint(buf[3])<<24
}

// ReadUint32 converts the first 4 bits in `b` into a uint32.
func ReadUint32(buf []byte) uint32 {
	return uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24
}

// ReadUint16 converts the first 2 bits in `b` into a uint16.
func ReadUint16(buf []byte) uint16 {
	return uint16(buf[0]) | uint16(buf[1])<<8
}

// WriteUint64 sets the first 8 bits in `b` to the value of `i`.
func WriteUint64(b []byte, u uint64) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
}

// WriteUintX64 sets the first 8 bits in `b` to the value of `i`.
func WriteUintX64(b []byte, u uint) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
}

// WriteUintX32 sets the first 4 bits in `b` to the value of `i`.
func WriteUintX32(b []byte, u uint) {
	b[0], b[1], b[2], b[3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
}

// WriteUint32 sets the first 4 bits in `b` to the value of `i`.
func WriteUint32(b []byte, u uint32) {
	b[0], b[1], b[2], b[3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
}

// WriteUint16 sets the first 2 bits in `b` to the value of `i`.
func WriteUint16(b []byte, u uint16) {
	b[0], b[1] = byte(u), byte(u>>_8)
}
