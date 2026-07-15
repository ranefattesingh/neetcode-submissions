func characterReplacement(s string, k int) int {
	m := make(map[byte]int)

	left := 0
	length := 0

	for right := 0; right < len(s); right++ {
		m[s[right]]++

		freq := 0
		for _, v := range m {
			freq = max(freq, v)
		}

		if (right - left + 1) - freq > k {
			m[s[left]]--
			left++
		}

		length = max(length, right - left + 1)
	}

	return length
}
