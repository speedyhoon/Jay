package jay

func SizeBools(length int) int {
	if length <= 0 {
		return 0
	}
	return (length + 7) >> 3
}

func WriteBools(y []byte, a []bool, length int) {
	for i, n := 0, 0; i*8 < length; i++ {
		n = i * 8
		switch length - n {
		case 1:
			y[i] = Bool1(a[n])
		case 2:
			y[i] = Bool2(a[n], a[n+1])
		case 3:
			y[i] = Bool3(a[n], a[n+1], a[n+2])
		case 4:
			y[i] = Bool4(a[n], a[n+1], a[n+2], a[n+3])
		case 5:
			y[i] = Bool5(a[n], a[n+1], a[n+2], a[n+3], a[n+4])
		case 6:
			y[i] = Bool6(a[n], a[n+1], a[n+2], a[n+3], a[n+4], a[n+5])
		case 7:
			y[i] = Bool7(a[n], a[n+1], a[n+2], a[n+3], a[n+4], a[n+5], a[n+6])
		default:
			y[i] = Bool8(a[n], a[n+1], a[n+2], a[n+3], a[n+4], a[n+5], a[n+6], a[n+7])
		}
	}
}

func ReadBools(y []byte, length int) (t []bool) {
	if length == 0 {
		return
	}
	t = make([]bool, length)
	for i, n := 0, 0; i*8 < length; i++ {
		n = i * 8
		switch length - n {
		case 1:
			t[n] = ReadBool1(y[i])
		case 2:
			t[n], t[n+1] = ReadBool2(y[i])
		case 3:
			t[n], t[n+1], t[n+2] = ReadBool3(y[i])
		case 4:
			t[n], t[n+1], t[n+2], t[n+3] = ReadBool4(y[i])
		case 5:
			t[n], t[n+1], t[n+2], t[n+3], t[n+4] = ReadBool5(y[i])
		case 6:
			t[n], t[n+1], t[n+2], t[n+3], t[n+4], t[n+5] = ReadBool6(y[i])
		case 7:
			t[n], t[n+1], t[n+2], t[n+3], t[n+4], t[n+5], t[n+6] = ReadBool7(y[i])
		default:
			t[n], t[n+1], t[n+2], t[n+3], t[n+4], t[n+5], t[n+6], t[n+7] = ReadBool8(y[i])
		}
	}
	return
}
