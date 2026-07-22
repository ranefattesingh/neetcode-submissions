func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxLength := 0
	left := 0
	right := 0
	for right = 0; right < len(s); right++ {
		if index, exists := m[s[right]]; exists {
			maxLength = max(maxLength, right - left)
			left = max(index + 1, left)
		}

		m[s[right]] = right
	}

	return max(maxLength, right - left)
}
