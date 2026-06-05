func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	result := 0
	for _, num := range nums {
		count := 0
		start := num
		for _, ok := m[start]; ok; _, ok = m[start] {
			count++
			start++
		}

		if count > result {
			result = count
		}
	}

	return result
}
