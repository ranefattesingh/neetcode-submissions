

func checkInclusion(s1 string, s2 string) bool {
	h1 := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		h1[s1[i]]++
	}

	need := len(h1)

	for i :=0 ; i < len(s2); i++ {
		h2 := make(map[byte]int)
		curr := 0
		for j := i; j < len(s2); j++ {
			h2[s2[j]]++
			if h1[s2[j]] < h2[s2[j]] {
				break
			}

			if h1[s2[j]] == h2[s2[j]] {
				curr++
			}

			if curr == need {
				return true
			}
		}

	}

	return false
}