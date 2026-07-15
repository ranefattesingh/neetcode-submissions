func characterReplacement(s string, k int) int {
	m := make(map[byte]int)

	left := 0
	length := 0
	maxFreq := 0

	for right := 0; right < len(s); right++ {
		m[s[right]]++

		maxFreq = max(maxFreq, m[s[right]])

		if (right - left + 1) - maxFreq > k {
			m[s[left]]--
			left++
		}

		length = max(length, right - left + 1)
	}

	return length
}
