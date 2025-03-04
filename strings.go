package jay

func SizeStrings8(s []string) (total int) {
	l := len(s)
	if l == 0 {
		return 1
	}

	for i := range s {
		total += len(s[i])
	}
	return 1 + l + total
}

func SizeStrings16(s []string) (total int) {
	l := len(s)
	if l == 0 {
		return 2
	}

	for i := range s {
		total += len(s[i])
	}
	return 2 + l*2 + total
}

func ReadStrings8(y []byte, s *[]string) (ok bool) {
	if y[0] == 0 {
		return true
	}

	l := len(y)
	qty := int(y[0])
	index := qty + 1
	if l < index {
		return false
	}

	total := index
	*s = make([]string, qty)
	for i := 0; i < qty; i++ {
		u := int(y[1+i])
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return false
		}

		(*s)[i] = string(y[index:total])
		index = total
	}

	return true
}

func ReadStrings8Err(y []byte, s *[]string) error {
	if y[0] == 0 {
		return nil
	}

	l := len(y)
	qty := int(y[0])
	index := qty + 1
	if l < index {
		return ErrUnexpectedEOB
	}

	total := index
	*s = make([]string, qty)
	for i := 0; i < qty; i++ {
		u := int(y[1+i])
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return ErrUnexpectedEOB
		}

		(*s)[i] = string(y[index:total])
		index = total
	}

	return nil
}

func ReadStrings8n(y []byte, s *[]string) (total int, ok bool) {
	if y[0] == 0 {
		return 0, true
	}

	l := len(y)
	qty := int(y[0])
	index := qty + 1
	if l < index {
		return 0, false
	}

	total = index
	*s = make([]string, qty)
	for i := 0; i < qty; i++ {
		u := int(y[1+i])
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return 0, false
		}

		(*s)[i] = string(y[index:total])
		index = total
	}

	return total, true
}

func ReadStrings8nErr(y []byte, s *[]string) (total int, err error) {
	if y[0] == 0 {
		return
	}

	l := len(y)
	qty := int(y[0])
	index := qty + 1
	if l < index {
		return 0, ErrUnexpectedEOB
	}

	total = index
	*s = make([]string, qty)
	for i := 0; i < qty; i++ {
		u := int(y[1+i])
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return 0, ErrUnexpectedEOB
		}

		(*s)[i] = string(y[index:total])
		index = total
	}

	return
}

func ReadStrings16(y []byte, s *[]string) (ok bool) {
	// ReadUint16
	qty := int(y[0]) | int(y[1])<<8
	if qty == 0 {
		return true
	}

	l := len(y)
	index := qty*2 + 2
	if l < index {
		return false
	}

	total := index
	*s = make([]string, qty)
	for i := 0; i < qty; i++ {
		// ReadUint16
		u := int(y[2+i*2]) | int(y[3+i*2])<<8
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return false
		}

		(*s)[i] = string(y[index:total])
		index = total
	}

	return true
}

func ReadStrings16Err(y []byte, s *[]string) error {
	// ReadUint16
	qty := int(y[0]) | int(y[1])<<8
	if qty == 0 {
		return nil
	}

	l := len(y)
	index := qty*2 + 2
	if l < index {
		return ErrUnexpectedEOB
	}

	total := index
	*s = make([]string, qty)
	for i := 0; i < qty; i++ {
		// ReadUint16
		u := int(y[2+i*2]) | int(y[3+i*2])<<8
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return ErrUnexpectedEOB
		}

		(*s)[i] = string(y[index:total])
		index = total
	}

	return nil
}

func ReadStrings16n(y []byte, s *[]string) (total int, ok bool) {
	// ReadUint16
	qty := int(y[0]) | int(y[1])<<8
	if qty == 0 {
		return 0, true
	}

	l := len(y)
	index := qty*2 + 2
	if l < index {
		return 0, false
	}

	total = index
	*s = make([]string, qty)
	for i := 0; i < qty; i++ {
		// ReadUint16
		u := int(y[2+i*2]) | int(y[3+i*2])<<8
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return 0, false
		}

		(*s)[i] = string(y[index:total])
		index = total
	}

	return total, true
}

func ReadStrings16nb(y []byte, s *[]string, at *int) (ok bool) {
	// ReadUint16
	qty := int(y[0]) | int(y[1])<<8
	if qty == 0 {
		return true
	}

	l := len(y)
	index := qty*2 + 2
	if l < index {
		return false
	}

	total := index
	*s = make([]string, qty)
	for i := 0; i < qty; i++ {
		// ReadUint16
		u := int(y[2+i*2]) | int(y[3+i*2])<<8
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return false
		}

		(*s)[i] = string(y[index:total])
		index = total
	}
	*at += total

	return true
}

func ReadStrings16nErr(y []byte, s *[]string) (total uint, err error) {
	// ReadUint16
	qty := uint(y[0]) | uint(y[1])<<8
	if qty == 0 {
		return
	}

	l := uint(len(y))
	index := qty*2 + 2
	if l < index {
		return 0, ErrUnexpectedEOB
	}

	total = index
	*s = make([]string, qty)
	for i := uint(0); i < qty; i++ {
		// ReadUint16
		u := uint(y[2+i*2]) | uint(y[3+i*2])<<8
		if u == 0 {
			continue
		}

		total += u
		if l < total {
			return 0, ErrUnexpectedEOB
		}

		(*s)[i] = string(y[index:total])
		index = total
	}

	return
}

func WriteStrings8(y []byte, s []string) {
	qty := len(s)
	if qty == 0 {
		return
	}

	y[0] = byte(qty)
	qty++

	var l, end int
	for i := range s {
		l = len(s[i])
		if l != 0 {
			y[1+i] = byte(l)
			end = qty + l
			copy(y[qty:end], s[i])
			qty = end
		}
	}
}

func WriteStrings16(y []byte, s []string) {
	qty := len(s)
	if qty == 0 {
		return
	}
	// WriteUint16
	y[0], y[1] = byte(qty), byte(qty>>8)
	const sizeBytes = 2
	qty = qty*sizeBytes + sizeBytes

	var l, end int
	for i := range s {
		l = len(s[i])
		if l != 0 {
			// WriteUint16
			y[2+i*sizeBytes], y[3+i*sizeBytes] = byte(l), byte(l>>8)
			end = qty + l
			copy(y[qty:end], s[i])
			qty = end
		}
	}
}
