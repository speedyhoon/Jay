//go:build 386 || amd64p32 || arm || armbe || mips || mips64p32 || mips64p32le || mipsle || ppc || riscv || s390 || sparc

// 32-bit only version. Complete list of GOARCH values: https://github.com/golang/go/blob/master/src/internal/syslist/syslist.go

package jay

// ReadInt for 32-bit systems.
func ReadInt(y []byte) (i, length int) {
	switch y[_0] {
	case _1:
		return int(int8(y[_1])), _2 // Convert to int8 to ensure negative integers are handled.
	case _2:
		return int(ReadInt16(y[_1:_3])), _3
	case _3:
		return ReadInt24(y[_1:_4]), _4
	case _4:
		return ReadIntX32(y[_1:_5]), _5
	}
	return
}
