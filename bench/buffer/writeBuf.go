package buffer

import "bytes"

const (
	maxUint8  = 1<<8 - 1  // 255
	maxUint16 = 1<<16 - 1 // 65535
	maxUint24 = 1<<24 - 1 // 16777215
	maxUint32 = 1<<32 - 1 // 4294967295
	maxUint40 = 1<<40 - 1 // 1099511627775
	maxUint48 = 1<<48 - 1 // 281474976710655
	maxUint56 = 1<<56 - 1 // 72057594037927935
	//maxUint64 = 1<<64 - 1 // 18446744073709551615
)

func WriteBufBool(b *bytes.Buffer, t bool) error {
	if t {
		return b.WriteByte(1)
	}

	return b.WriteByte(0)
}

// WriteBufBools ...
func WriteBufBools(b *bytes.Buffer, t []bool) (err error) {
	l := len(t)
	if l == 0 {
		return b.WriteByte(0)
	}

	err = b.WriteByte(byte(l))
	if err != nil {
		return
	}

	for i := range t {
		err = WriteBufBool(b, t[i])
		if err != nil {
			return
		}
	}

	return nil
}

func WriteBufString(b *bytes.Buffer, s string) (err error) {
	if l := len(s); l >= 1 {
		err = b.WriteByte(byte(l))
		if err != nil {
			return
		}
		_, err = b.WriteString(s)
		return
	}

	return b.WriteByte(0)
}

func WriteBufUint(b *bytes.Buffer, i uint) (err error) {
	switch {
	case i == 0:
		return b.WriteByte(0)
	case i <= maxUint8:
		_, err = b.Write([]byte{1, byte(i)})
	case i <= maxUint16:
		_, err = b.Write([]byte{2, byte(i), byte(i >> 8)})
	case i <= maxUint24:
		_, err = b.Write([]byte{3, byte(i), byte(i >> 8), byte(i >> 16)})
	case i <= maxUint32:
		_, err = b.Write([]byte{4, byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)})
	case i <= maxUint40:
		_, err = b.Write([]byte{5, byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32)})
	case i <= maxUint48:
		_, err = b.Write([]byte{6, byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32), byte(i >> 40)})
	case i <= maxUint56:
		_, err = b.Write([]byte{7, byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32), byte(i >> 40), byte(i >> 48)})
	default:
		_, err = b.Write([]byte{8, byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32), byte(i >> 40), byte(i >> 48), byte(i >> 56)})
	}
	return
}

func WriteBufUint64(b *bytes.Buffer, i uint64) (err error) {
	_, err = b.Write([]byte{8, byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32), byte(i >> 40), byte(i >> 48), byte(i >> 56)})
	return
}

func WriteBufUint64s(b *bytes.Buffer, u []uint64) (err error) {
	l := len(u)
	if l == 0 {
		return b.WriteByte(0)
	}

	err = b.WriteByte(byte(l))
	if err != nil {
		return
	}

	for i := range u {
		err = WriteBufUint64(b, u[i])
		if err != nil {
			return
		}
	}
	return nil
}

func WriteBufUints(b *bytes.Buffer, u []uint) (err error) {
	l := len(u)
	if l == 0 {
		return b.WriteByte(0)
	}

	err = b.WriteByte(byte(l))
	if err != nil {
		return
	}

	for i := range u {
		err = WriteBufUint(b, u[i])
		if err != nil {
			return
		}
	}
	return nil
}

func WriteBufUint32s(b *bytes.Buffer, u []uint32) (err error) {
	l := len(u)
	if l == 0 {
		return b.WriteByte(0)
	}

	err = b.WriteByte(byte(l))
	if err != nil {
		return
	}

	for i := range u {
		err = WriteBufUint32(b, u[i])
		if err != nil {
			return
		}
	}
	return nil
}

func WriteBufUint16s(b *bytes.Buffer, u []uint16) (err error) {
	l := len(u)
	if l == 0 {
		return b.WriteByte(0)
	}

	err = b.WriteByte(byte(l))
	if err != nil {
		return
	}

	for i := range u {
		err = WriteBufUint16(b, u[i])
		if err != nil {
			return
		}
	}
	return nil
}

func WriteBufUint32(b *bytes.Buffer, i uint32) (err error) {
	_, err = b.Write([]byte{4, byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)})
	return
}

func WriteBufUint16(b *bytes.Buffer, i uint16) (err error) {
	_, err = b.Write([]byte{2, byte(i), byte(i >> 8)})
	return
}

func WriteBufStrings(b *bytes.Buffer, s []string) (err error) {
	l := len(s)
	if l == 0 {
		return b.WriteByte(0)
	}

	err = b.WriteByte(byte(l))
	if err != nil {
		return
	}

	for i := range s {
		err = WriteBufString(b, s[i])
		if err != nil {
			return
		}
	}
	return
}
