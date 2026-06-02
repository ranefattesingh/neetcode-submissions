func hasDuplicate(nums []int) bool {
    m := make(map[int]struct{})

	for _, v := range nums {
		_, exists := m[v]
		if exists {
			return true
		}

		m[v] = struct{}{}
	}

	return false
}
