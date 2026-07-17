func checkInclusion(s1 string, s2 string) bool {
	// Edge case when s1 > s2
	if len(s1) > len(s2) {
		return false
	}

	// Frequency maps for s1 and current window
	s1Count := make(map[byte]int)
	windowCount := make(map[byte]int)

	// Initialize first window
	for right := 0; right < len(s1); right++ {
		s1Count[s1[right]]++
		windowCount[s2[right]]++
	}

	// Check first window
	if equal(s1Count, windowCount) {
		return true
	}

	// Slide the window
	left := 0
	for right := len(s1); right < len(s2); right++ {
		// Add right char to window
		windowCount[s2[right]]++

		// Remove left char
		windowCount[s2[left]]--
		if windowCount[s2[left]] == 0 {
			delete(windowCount, s2[left])
		}
		left++

		if equal(s1Count, windowCount) {
			return true
		}
	}

	return false
}

func equal(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if b[k] != v {
			return false
		}
	}

	return true
}