package jay

// Bit shift masks.
const _24, _40, _48, _56 = 24, 40, 48, 56

// ReadInt64 converts the first 8 bits in `b` into an int64.
func ReadInt64(y []byte) int64 {
	return int64(y[0]) | int64(y[1])<<_8 | int64(y[2])<<_16 | int64(y[3])<<_24 |
		int64(y[4])<<_32 | int64(y[5])<<_40 | int64(y[6])<<_48 | int64(y[7])<<_56
}

// ReadIntX64 converts the first 8 bits in `b` into an int.
func ReadIntX64(y []byte) int {
	return int(y[0]) | int(y[1])<<_8 | int(y[2])<<_16 | int(y[3])<<_24 |
		int(y[4])<<_32 | int(y[5])<<_40 | int(y[6])<<_48 | int(y[7])<<_56
}

// ReadIntX32 converts the first 6 bits in `b` into an int.
func ReadIntX32(y []byte) int {
	return int(int32(y[0]) | int32(y[1])<<_8 | int32(y[2])<<_16 | int32(y[3])<<_24)
}

// ReadInt32 converts the first 4 bits in `b` into an int32.
func ReadInt32(y []byte) int32 {
	return int32(y[0]) | int32(y[1])<<_8 | int32(y[2])<<_16 | int32(y[3])<<_24
}

// ReadInt16 converts the first 2 bits in `b` into an int16.
func ReadInt16(y []byte) int16 {
	return int16(y[0]) | int16(y[1])<<_8
}

// WriteInt64 sets the first 8 bits in `b` to the value of `i`.
func WriteInt64(y []byte, i int64) {
	y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48), byte(i>>_56)
}

// WriteIntX64 sets the first 8 bits in `b` to the value of `i`.
func WriteIntX64(y []byte, i int) {
	y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48), byte(i>>_56)
}

// WriteIntX32 sets the first 4 bits in `b` to the value of `i`.
func WriteIntX32(y []byte, i int) {
	y[0], y[1], y[2], y[3] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24)
}

// WriteInt32 sets the first 4 bits in `b` to the value of `i`.
func WriteInt32(y []byte, i int32) {
	y[0], y[1], y[2], y[3] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24)
}

// WriteInt16 sets the first 2 bits in `b` to the value of `i`.
func WriteInt16(y []byte, i int16) {
	y[0], y[1] = byte(i), byte(i>>_8)
}
