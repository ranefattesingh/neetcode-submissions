func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count := [26]int{}
	s2Count := [26]int{}

	right := 0
	for right < len(s1) {
		s1Count[s1[right]-'a']++
		s2Count[s2[right]-'a']++
		right++
	}

	matches := 0
	for i := range 26 {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}

	left := 0
	for right < len(s2) {
		if matches == 26 {
			return true
		}

		index := s2[right] - 'a'
		s2Count[index]++
		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]+1 == s2Count[index] {
			matches--
		}

		index = s2[left] - 'a'
		s2Count[index]--
		if s1Count[index] == s2Count[index] {
			matches++
		} else if s1Count[index]-1 == s2Count[index] {
			matches--
		}

		left++
		right++ 
	}

	return matches == 26
}