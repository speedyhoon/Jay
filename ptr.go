package jay

func LenStr(s *string) int {
	if s == nil {
		return 0
	}

	return len(*s)
}
