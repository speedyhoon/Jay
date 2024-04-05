package generate

func RemoveDuplicates(s []string) []string {
	// Loop through the first item through to the second last item.
	for i := 0; i < len(s)-1; i++ {
		// Loop from item i+1 through to the last item.
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				s = Remove(s, j)
				j--
			}
		}
	}
	return s
}
