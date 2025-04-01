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
