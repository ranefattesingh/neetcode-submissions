func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return  false
	}

	counts := make([]int, 26)

	for i  := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for i  := 0; i < len(s); i++ {
		if counts[s[i]-'a'] != 0 {
			return false
		}

		if counts[t[i]-'a'] != 0 {
			return false
		}
	}

	return true
}
