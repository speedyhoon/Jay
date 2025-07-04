package buffer

import (
	"bytes"
	"log"
)

func main() {
	u := uint(654)

	buf := bytes.NewBuffer(nil)

	//binary.LittleEndian.PutUint32()
	buf.WriteByte(byte(u))
	buf.WriteByte(byte(u >> 8))
	buf.WriteByte(byte(u >> 16))
	buf.WriteByte(byte(u >> 24))

	//PutUint32(buf, u)

	log.Println(buf.String())
}

func WriteUint(buf *bytes.Buffer, u uint) {
	buf.WriteByte(byte(u))
	buf.WriteByte(byte(u >> 8))
	buf.WriteByte(byte(u >> 16))
	buf.WriteByte(byte(u >> 24))
}

func WriteUint2(buf *bytes.Buffer, u uint) {

	buf.WriteByte(byte(u))
	buf.WriteByte(byte(u >> 8))
	buf.WriteByte(byte(u >> 16))
	buf.WriteByte(byte(u >> 24))
}

func PutUint32(b []byte, v uint) {
	if len(b) < 4 {
		return
	}
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func WriteBytes(buf *bytes.Buffer, y ...byte) {
	buf.Write(y)
}

func WriteUintBytes(buf *bytes.Buffer, u uint) int {
	if u > maxUint56 {
		WriteBytes(buf, 8, byte(u), byte(u>>8), byte(u>>16), byte(u>>24), byte(u>>32), byte(u>>40), byte(u>>48), byte(u>>56))
		//buf.WriteByte(byte(u))
		//buf.WriteByte(byte(u >> 8))
		//buf.WriteByte(byte(u >> 16))
		//buf.WriteByte(byte(u >> 24))
		//buf.WriteByte(byte(u >> 32))
		//buf.WriteByte(byte(u >> 40))
		//buf.WriteByte(byte(u >> 48))
		//buf.WriteByte(byte(u >> 56))
		return 9
	}

	if u > maxUint48 {
		WriteBytes(buf, 7, byte(u), byte(u>>8), byte(u>>16), byte(u>>24), byte(u>>32), byte(u>>40), byte(u>>48))
		return 8
	}

	if u > maxUint40 {
		WriteBytes(buf, 6, byte(u), byte(u>>8), byte(u>>16), byte(u>>24), byte(u>>32), byte(u>>40))
		return 7
	}

	if u > maxUint32 {
		WriteBytes(buf, 5, byte(u), byte(u>>8), byte(u>>16), byte(u>>24), byte(u>>32))
		return 6
	}

	if u > maxUint24 {
		WriteBytes(buf, 4, byte(u), byte(u>>8), byte(u>>16), byte(u>>24))
		return 5
	}

	if u > maxUint16 {
		WriteBytes(buf, 3, byte(u), byte(u>>8), byte(u>>16))
		return 4
	}

	if u > maxUint8 {
		WriteBytes(buf, 2, byte(u), byte(u>>8))
		return 3
	}

	if u == 0 {
		buf.WriteByte(0)
		return 1
	}

	WriteBytes(buf, 1, byte(u))
	return 2

	/*
		switch {
		case u == 0:
			buf.WriteByte(0)
			return 1
		case u <= maxUint8:
			buf.Write([]byte{1, byte(u)})
			return 2
		case u <= maxUint16:
			buf.WriteByte(2)
			buf.WriteByte(byte(u))
			buf.WriteByte(byte(u >> 8))
			return 3
		case u <= maxUint24:
			buf.WriteByte(3)
			buf.WriteByte(byte(u))
			buf.WriteByte(byte(u >> 8))
			buf.WriteByte(byte(u >> 16))
			return 4
		case u <= maxUint32:
			buf.WriteByte(4)
			buf.WriteByte(byte(u))
			buf.WriteByte(byte(u >> 8))
			buf.WriteByte(byte(u >> 16))
			buf.WriteByte(byte(u >> 24))
			return 5
		default:
			buf.WriteByte(9)
			buf.WriteByte(byte(u))
			buf.WriteByte(byte(u >> 8))
			buf.WriteByte(byte(u >> 16))
			buf.WriteByte(byte(u >> 24))
			return 9
		}*/
}

func WriteUintBytes2(buf *bytes.Buffer, u uint) int {
	if u > maxUint56 {
		buf.Write([]byte{8, byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24), byte(u >> 32), byte(u >> 40), byte(u >> 48), byte(u >> 56)})
		return 9
	}

	if u > maxUint48 {
		buf.Write([]byte{7, byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24), byte(u >> 32), byte(u >> 40), byte(u >> 48)})
		return 8
	}

	if u > maxUint40 {
		buf.Write([]byte{6, byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24), byte(u >> 32), byte(u >> 40)})
		return 7
	}

	if u > maxUint32 {
		buf.Write([]byte{5, byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24), byte(u >> 32)})
		return 6
	}

	if u > maxUint24 {
		buf.Write([]byte{4, byte(u), byte(u >> 8), byte(u >> 16), byte(u >> 24)})
		return 5
	}

	if u > maxUint16 {
		buf.Write([]byte{3, byte(u), byte(u >> 8), byte(u >> 16)})
		return 4
	}

	if u > maxUint8 {
		buf.Write([]byte{2, byte(u), byte(u >> 8)})
		return 3
	}

	if u == 0 {
		buf.WriteByte(0)
		return 1
	}

	buf.Write([]byte{1, byte(u)})
	return 2
}
func WriteUintBytes3(buf *bytes.Buffer, u uint) int {
	if u > maxUint56 {
		buf.WriteByte(8)
		buf.WriteByte(byte(u))
		buf.WriteByte(byte(u >> 8))
		buf.WriteByte(byte(u >> 16))
		buf.WriteByte(byte(u >> 24))
		buf.WriteByte(byte(u >> 32))
		buf.WriteByte(byte(u >> 40))
		buf.WriteByte(byte(u >> 48))
		buf.WriteByte(byte(u >> 56))
		return 9
	}

	if u > maxUint48 {
		buf.WriteByte(7)
		buf.WriteByte(byte(u))
		buf.WriteByte(byte(u >> 8))
		buf.WriteByte(byte(u >> 16))
		buf.WriteByte(byte(u >> 24))
		buf.WriteByte(byte(u >> 32))
		buf.WriteByte(byte(u >> 40))
		buf.WriteByte(byte(u >> 48))
		return 8
	}

	if u > maxUint40 {
		buf.WriteByte(6)
		buf.WriteByte(byte(u))
		buf.WriteByte(byte(u >> 8))
		buf.WriteByte(byte(u >> 16))
		buf.WriteByte(byte(u >> 24))
		buf.WriteByte(byte(u >> 32))
		buf.WriteByte(byte(u >> 40))
		return 7
	}

	if u > maxUint32 {
		buf.WriteByte(5)
		buf.WriteByte(byte(u))
		buf.WriteByte(byte(u >> 8))
		buf.WriteByte(byte(u >> 16))
		buf.WriteByte(byte(u >> 24))
		buf.WriteByte(byte(u >> 32))
		return 6
	}

	if u > maxUint24 {
		buf.WriteByte(4)
		buf.WriteByte(byte(u))
		buf.WriteByte(byte(u >> 8))
		buf.WriteByte(byte(u >> 16))
		buf.WriteByte(byte(u >> 24))
		return 5
	}

	if u > maxUint16 {
		buf.WriteByte(3)
		buf.WriteByte(byte(u))
		buf.WriteByte(byte(u >> 8))
		buf.WriteByte(byte(u >> 16))
		return 4
	}

	if u > maxUint8 {
		buf.WriteByte(2)
		buf.WriteByte(byte(u))
		buf.WriteByte(byte(u >> 8))
		return 3
	}

	if u == 0 {
		buf.WriteByte(0)
		return 1
	}

	buf.WriteByte(1)
	buf.WriteByte(byte(u))
	return 2
}
