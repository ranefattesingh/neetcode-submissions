func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Count, s2Count [26]int
	for i, v := range s1 {
		s1Count[v-'a']++
		s2Count[s2[i]-'a']++
	}

	for right := len(s1); right < len(s2); right++ {
		if s1Count == s2Count {
			return true
		}

		s2Count[s2[right]-'a']++
		s2Count[s2[right-len(s1)]-'a']--
	}

	return s1Count == s2Count
}