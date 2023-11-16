package jay

const _128, _64, _32, _16, _8, _4, _2, _1 = 128, 64, 32, 16, 8, 4, 2, 1
const _24, _40, _48, _56 = 24, 40, 48, 56

func Bool1(a bool) (y byte) {
	if a {
		return _128
	}
	return
}

func Bool2(a, b bool) (y byte) {
	if a {
		y = _128
	}
	if b {
		y |= _64
	}
	return
}

func Bool3(a, b, c bool) (y byte) {
	if a {
		y = _128
	}
	if b {
		y |= _64
	}
	if c {
		y |= _32
	}
	return
}

func Bool4(a, b, c, d bool) (y byte) {
	if a {
		y = _128
	}
	if b {
		y |= _64
	}
	if c {
		y |= _32
	}
	if d {
		y |= _16
	}
	return
}

func Bool5(a, b, c, d, e bool) (y byte) {
	if a {
		y = _128
	}
	if b {
		y |= _64
	}
	if c {
		y |= _32
	}
	if d {
		y |= _16
	}
	if e {
		y |= _8
	}
	return
}

func Bool6(a, b, c, d, e, f bool) (y byte) {
	if a {
		y = _128
	}
	if b {
		y |= _64
	}
	if c {
		y |= _32
	}
	if d {
		y |= _16
	}
	if e {
		y |= _8
	}
	if f {
		y |= _4
	}
	return
}

func Bool7(a, b, c, d, e, f, g bool) (y byte) {
	if a {
		y = _128
	}
	if b {
		y |= _64
	}
	if c {
		y |= _32
	}
	if d {
		y |= _16
	}
	if e {
		y |= _8
	}
	if f {
		y |= _4
	}
	if g {
		y |= _2
	}
	return
}

func Bool8(a, b, c, d, e, f, g, h bool) (y byte) {
	if a {
		y = _128
	}
	if b {
		y |= _64
	}
	if c {
		y |= _32
	}
	if d {
		y |= _16
	}
	if e {
		y |= _8
	}
	if f {
		y |= _4
	}
	if g {
		y |= _2
	}
	if h {
		y |= _1
	}
	return
}

func ReadBool1(y byte) (a bool) {
	return y >= _128
}

func ReadBool2(y byte) (a, b bool) {
	return y >= _128, y&_64 == _64
}

func ReadBool3(y byte) (a, b, c bool) {
	return y >= _128, y&_64 == _64, y&_32 == _32
}

func ReadBool4(y byte) (a, b, c, d bool) {
	return y >= _128, y&_64 == _64, y&_32 == _32, y&_16 == _16
}

func ReadBool5(y byte) (a, b, c, d, e bool) {
	return y >= _128, y&_64 == _64, y&_32 == _32, y&_16 == _16, y&_8 == _8
}

func ReadBool6(y byte) (a, b, c, d, e, f bool) {
	return y >= _128, y&_64 == _64, y&_32 == _32, y&_16 == _16, y&_8 == _8, y&_4 == _4
}

func ReadBool7(y byte) (a, b, c, d, e, f, g bool) {
	return y >= _128, y&_64 == _64, y&_32 == _32, y&_16 == _16, y&_8 == _8, y&_4 == _4, y&_2 == _2
}

func ReadBool8(y byte) (a, b, c, d, e, f, g, h bool) {
	return y >= _128, y&_64 == _64, y&_32 == _32, y&_16 == _16, y&_8 == _8, y&_4 == _4, y&_2 == _2, y&_1 == _1
}
