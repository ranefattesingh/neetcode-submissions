func findDuplicate(nums []int) int {
	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] < 0 {
			return abs(num)
		}

		nums[index] *= -1
	}

	return -1
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
