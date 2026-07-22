func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	maxLength := 0
	left := 0
	right := 0
	for right = 0; right < len(s); right++ {
		for exists := m[s[right]]; exists; exists = m[s[right]] {
			delete(m, s[left])
			maxLength = max(maxLength, right - left)

			left++
		}

		m[s[right]] = true
	}

	return max(maxLength, right - left)
}
