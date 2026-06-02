func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	n := make(map[rune]int)

	for _, v := range t {
		if _, ok := n[v]; !ok {
			n[v] = 1
		} else {
			n[v]++
		}
	}

	if len(m) != len(n) {
		return false
	}

	for k, v := range m {
		nv, exists := n[k]
		if !exists || nv != v {
			return false
		}
	}


	return true
}
