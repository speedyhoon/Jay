package jay

// Read64Int ...
func Read64Int(b []byte) int {
	return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24 |
		int(b[4])<<_32 | int(b[5])<<_40 | int(b[6])<<_48 | int(b[7])<<_56
}

// Read32Int ...
func Read32Int(b []byte) int {
	return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24
}

// Write64Int ...
func Write64Int(b []byte, i int) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48), byte(i>>_56)
}

// Write32Int ...
func Write32Int(b []byte, i int) {
	b[0], b[1], b[2], b[3] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24)
}
