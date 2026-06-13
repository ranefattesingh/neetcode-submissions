func hasDuplicate(nums []int) bool {
	hashMap := make(map[int]struct{})

	for _, v := range nums {
		_, ok := hashMap[v]
		if ok {
			return true
		}

		hashMap[v] = struct{}{}
	}

	return false
}
