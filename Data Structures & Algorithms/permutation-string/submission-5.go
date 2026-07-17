func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Count, windowCount [26]int

	// Build frequency arrays
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		windowCount[s2[i]-'a']++
	}

	// Check first window
	if s1Count == windowCount {
		return true
	}

	// Slide the window
	left := 0
	for right := len(s1); right < len(s2); right++ {
		windowCount[s2[right]-'a']++
		windowCount[s2[left]-'a']--
		left++

		if s1Count == windowCount {
			return true
		}
	}

	return false
}