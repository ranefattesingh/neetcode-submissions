func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for next, v := range nums {
		prev, ok := hashMap[v]
		if ok {
			return []int{prev, next}
		}

		hashMap[target - v] = next
	}

	return nil
}
