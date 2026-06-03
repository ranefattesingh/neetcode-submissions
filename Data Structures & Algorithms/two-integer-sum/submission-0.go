func twoSum(nums []int, target int) []int {
    store := make(map[int]int)

	for index, value := range nums {
		prevIndex, exists := store[value]
		if exists {
			return []int{prevIndex, index}
		}

		store[target - value] = index
	}

	return nil
}
