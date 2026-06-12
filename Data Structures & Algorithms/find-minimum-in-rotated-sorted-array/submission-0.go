func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	globalMin := nums[0]

	for l <= r {
		m := l + (r - l) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m - 1
		}

		globalMin = min(globalMin, nums[m])
	}

	return globalMin
}