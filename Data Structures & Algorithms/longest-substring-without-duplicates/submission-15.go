func lengthOfLongestSubstring(s string) int {
	// zzxxyxyzxyz
	// z:2 x:3 y1 

	left := 0
	right := 0

	m := make(map[byte]struct{})
	maximum := 0
	currentMaximum := 0

	for ; right < len(s); right++ {
		for _, ok := m[s[right]]; ok; _, ok = m[s[right]] {
			delete(m, s[left])
			left++
		}

		currentMaximum = right - left + 1
		maximum = max(maximum, currentMaximum)
		m[s[right]] = struct{}{}
	}


	return maximum
}