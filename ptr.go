package jay

// LenStr returns the length of the string pointer.
func LenStr(s *string) int {
	if s == nil {
		return 0
	}

	return len(*s)
}
