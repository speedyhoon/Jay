package jay

const maxUint8 = 255

func SizeStrings8(s []string) (total int) {
	l := len(s) & maxUint8 // Truncate lengths > 255.
	if l == _0 {
		return
	}

	for i := _0; i < l; i++ {
		total += len(s[i]) & maxUint8
	}
	return l + total
}

func SizeStringsArray(s []string, l int) (total int) {
	for i := _0; i < l; i++ {
		total += len(s[i]) & maxUint8
	}
	return l + total
}

func ReadStrings8nbX(y []byte, s *[]string, qty uint8, at *int) (ok bool) {
	if qty == _0 {
		return true
	}

	l, index := len(y), int(qty)
	if l < index {
		return false
	}

	end := index
	*s = make([]string, qty)
	for i := uint8(_0); i < qty; i, index = i+1, end {
		end += int(y[i])
		if index == end {
			continue
		}

		if l < end {
			return false
		}

		(*s)[i] = string(y[index:end])
	}
	*at += end

	return true
}

func ReadStrings8Err(y []byte, s *[]string, qty uint8) (err error) {
	if qty == 0 {
		return
	}

	l, at := len(y), int(qty)
	if l < at {
		return ErrUnexpectedEOB
	}

	*s = make([]string, qty)
	for i, end := uint8(0), at; i < qty; i, at = i+1, end {
		end += int(y[i])
		if at == end {
			continue
		}

		if l < end {
			return ErrUnexpectedEOB
		}

		(*s)[i] = string(y[at:end])
	}

	return
}

func ReadStrings8nbXt(y []byte, s *[]string, qty uint8, at int) (end int, ok bool) {
	if qty == 0 {
		return at, true
	}

	l, index := len(y), int(qty)
	if l < index {
		return at, false
	}

	end = index
	*s = make([]string, qty)
	for i := uint8(0); i < qty; i, index = i+1, end {
		end += int(y[i])
		if index == end {
			continue
		}

		if l < end {
			return at, false
		}
		(*s)[i] = string(y[index:end])
	}

	return at + end, true
}

func ReadStrings8Ok(y []byte, s *[]string, qty uint8) (ok bool) {
	if qty == 0 {
		return true
	}

	l, at := len(y), int(qty)
	if l < at {
		return false
	}

	*s = make([]string, qty)
	for i, end := uint8(0), at; i < qty; i, at = i+1, end {
		end += int(y[i])
		if at == end {
			continue
		}

		if l < end {
			return false
		}

		(*s)[i] = string(y[at:end])
	}

	return true
}

// WriteStrings8Req is tagged as required - so should never be 0 or null
func WriteStrings8Req(y /*, qty*/ []byte, s []string, qty uint8) {
	length := len(s)
	if length > maxUint8 {
		panic("dif size lengths")
	}

	for i, at, end := 0, int(qty), int(qty); i < int(qty); i++ {
		y[i] = byte(len(s[i]))

		length = len(s[i])
		if length > maxUint8 {
			panic("dif rune lengths")
		}

		if y[i] != 0 {
			end += int(y[i])
			copy(y[at:end], s[i])
			at = end
		}
	}
}

// WriteStrings8 no tag (default) - may have nil or zero length slices.
func WriteStrings8(y, qty []byte, s []string) {
	length := len(s)
	if length > maxUint8 {
		panic("dif size lengths")
	}

	l := len(s) & maxUint8
	if l == 0 {
		return
	}
	qty[0] = byte(l)

	for i, at, end := 0, l, l; i < l; i++ {
		y[i] = byte(len(s[i]))

		length = len(s[i])
		if length > maxUint8 {
			panic("dif rune lengths")
		}

		if y[i] != 0 {
			end += int(y[i])
			copy(y[at:end], s[i])
			at = end
		}
	}
}

func WriteStringsArray(y []byte, l uint8, s []string) {
	for i, at, end := uint8(0), uint(l), uint(l); i < l; i++ {
		length := len(s[i])
		if length == 0 {
			continue
		}

		if length > maxUint8 {
			panic("dif rune lengths")
		}
		y[i] = byte(len(s[i]))

		end += uint(y[i])
		copy(y[at:end], s[i])
		at = end
	}
}

func ReadStringsArrayErr(y []byte, s []string, qty uint8) (err error) {
	for i, at, end, l := uint8(0), uint(qty), uint(qty), uint(len(y)); i < qty; i, at = i+1, end {
		end += uint(y[i])
		if at == end {
			continue
		}

		if l < end {
			return ErrUnexpectedEOB
		}

		s[i] = string(y[at:end])
	}

	return
}

// ReadStringsArrayAtOk decodes a []byte into []string and increments atPos with the overall length.
func ReadStringsArrayAtOk(y []byte, s []string, qty uint8, atPos *uint) (ok bool) {
	end := uint(qty)
	for i, at, l := uint8(0), uint(qty), uint(len(y)); i < qty; i, at = i+1, end {
		end += uint(y[i])
		if at == end {
			continue
		}

		if l < end {
			return false
		}

		s[i] = string(y[at:end])
	}

	*atPos += end
	return true
}
