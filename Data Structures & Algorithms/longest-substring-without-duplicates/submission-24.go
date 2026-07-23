func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maximum := 0

	left := 0
	for right := 0; right < len(s); right++ {
		jump, ok := m[s[right]] 
		if ok {
			left = max(jump + 1, left)
		}

		m[s[right]] = right
		maximum = max(maximum, right - left + 1)
	}

	return maximum
}