package jay

// WriteUint64 ...
func WriteUint64(b []byte, u uint64) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
}

// WriteUintArch64 ...
func WriteUintArch64(b []byte, u uint) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24), byte(u>>_32), byte(u>>_40), byte(u>>_48), byte(u>>_56)
}

// WriteUint32 ...
func WriteUint32(b []byte, u uint32) {
	b[0], b[1], b[2], b[3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
}

// WriteUintArch32 ...
func WriteUintArch32(b []byte, u uint) {
	b[0], b[1], b[2], b[3] = byte(u), byte(u>>_8), byte(u>>_16), byte(u>>_24)
}

// WriteUint16 ...
func WriteUint16(b []byte, u uint16) {
	b[0], b[1] = byte(u), byte(u>>_8)
}
