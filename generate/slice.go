package generate

func RemoveDuplicates(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				s = Remove(s, j)
				j--
			}
		}
	}
	return s
}
