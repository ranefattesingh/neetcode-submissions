func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int)
	for _, v := range s {
		count, _ := mapS[v]
		mapS[v] = count + 1
	}

	mapT := make(map[rune]int)
	for _, v := range t {
		count, _ := mapT[v]
		mapT[v] = count + 1
	}

	for key, sCount := range mapS {
		tCount, _ := mapT[key]
		if tCount != sCount {
			return false
		}
	}


	return true
}
