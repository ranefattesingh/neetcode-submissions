func sortBytes(b []byte) {
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
}

func checkInclusion(s1 string, s2 string) bool {
	b1 := []byte(s1)
	sortBytes(b1)
	s1 = string(b1)

	left := 0
	for right := 0; right < len(s2); right++ {
		if right-left+1 == len(s1) {
			b := []byte(s2[left : right+1])
			sortBytes(b)
			left++

			if string(b) == s1 {
				return true
			}
		}
	}

	return false
}