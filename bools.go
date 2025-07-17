package jay

func SizeBools(length int) int {
	if length <= _0 {
		return _0
	}
	return (length + _7) >> _3
}

func SizeBools8(length uint8) int {
	return (int(length) + _7) >> _3
}

func WriteBools(y []byte, a []bool, length int) {
	for i, n := _0, _0; n < length; i, n = i+_1, n+_8 {
		switch length - n {
		case _1:
			y[i] = Bool1(a[n])
		case _2:
			y[i] = Bool2(a[n], a[n+_1])
		case _3:
			y[i] = Bool3(a[n], a[n+_1], a[n+_2])
		case _4:
			y[i] = Bool4(a[n], a[n+_1], a[n+_2], a[n+_3])
		case _5:
			y[i] = Bool5(a[n], a[n+_1], a[n+_2], a[n+_3], a[n+_4])
		case _6:
			y[i] = Bool6(a[n], a[n+_1], a[n+_2], a[n+_3], a[n+_4], a[n+_5])
		case _7:
			y[i] = Bool7(a[n], a[n+_1], a[n+_2], a[n+_3], a[n+_4], a[n+_5], a[n+_6])
		default:
			y[i] = Bool8(a[n], a[n+_1], a[n+_2], a[n+_3], a[n+_4], a[n+_5], a[n+_6], a[n+_7])
		}
	}
}

func WriteBools8(y []byte, a []bool, length uint8) {
	var i, n uint8
	for ; n < length; i, n = i+_1, n+_8 {
		switch length - n {
		case _1:
			y[i] = Bool1(a[n])
		case _2:
			y[i] = Bool2(a[n], a[n+_1])
		case _3:
			y[i] = Bool3(a[n], a[n+_1], a[n+_2])
		case _4:
			y[i] = Bool4(a[n], a[n+_1], a[n+_2], a[n+_3])
		case _5:
			y[i] = Bool5(a[n], a[n+_1], a[n+_2], a[n+_3], a[n+_4])
		case _6:
			y[i] = Bool6(a[n], a[n+_1], a[n+_2], a[n+_3], a[n+_4], a[n+_5])
		case _7:
			y[i] = Bool7(a[n], a[n+_1], a[n+_2], a[n+_3], a[n+_4], a[n+_5], a[n+_6])
		default:
			y[i] = Bool8(a[n], a[n+_1], a[n+_2], a[n+_3], a[n+_4], a[n+_5], a[n+_6], a[n+_7])
		}
	}
}

func ReadBools(y []byte, length int) (t []bool) {
	if length == _0 {
		return
	}
	t = make([]bool, length)
	for i, n := _0, _0; n < length; i, n = i+_1, n+_8 {
		switch length - n {
		case _1:
			t[n] = ReadBool1(y[i])
		case _2:
			t[n], t[n+_1] = ReadBool2(y[i])
		case _3:
			t[n], t[n+_1], t[n+_2] = ReadBool3(y[i])
		case _4:
			t[n], t[n+_1], t[n+_2], t[n+_3] = ReadBool4(y[i])
		case _5:
			t[n], t[n+_1], t[n+_2], t[n+_3], t[n+_4] = ReadBool5(y[i])
		case _6:
			t[n], t[n+_1], t[n+_2], t[n+_3], t[n+_4], t[n+_5] = ReadBool6(y[i])
		case _7:
			t[n], t[n+_1], t[n+_2], t[n+_3], t[n+_4], t[n+_5], t[n+_6] = ReadBool7(y[i])
		default:
			t[n], t[n+_1], t[n+_2], t[n+_3], t[n+_4], t[n+_5], t[n+_6], t[n+_7] = ReadBool8(y[i])
		}
	}
	return
}

func ReadBools8(y []byte, length uint8) (t []bool) {
	if length == _0 {
		return
	}
	t = make([]bool, length)
	for i, n := _0, uint8(_0); n < length; i, n = i+_1, n+_8 {
		switch length - n {
		case _1:
			t[n] = ReadBool1(y[i])
		case _2:
			t[n], t[n+_1] = ReadBool2(y[i])
		case _3:
			t[n], t[n+_1], t[n+_2] = ReadBool3(y[i])
		case _4:
			t[n], t[n+_1], t[n+_2], t[n+_3] = ReadBool4(y[i])
		case _5:
			t[n], t[n+_1], t[n+_2], t[n+_3], t[n+_4] = ReadBool5(y[i])
		case _6:
			t[n], t[n+_1], t[n+_2], t[n+_3], t[n+_4], t[n+_5] = ReadBool6(y[i])
		case _7:
			t[n], t[n+_1], t[n+_2], t[n+_3], t[n+_4], t[n+_5], t[n+_6] = ReadBool7(y[i])
		default:
			t[n], t[n+_1], t[n+_2], t[n+_3], t[n+_4], t[n+_5], t[n+_6], t[n+_7] = ReadBool8(y[i])
		}
	}
	return
}
