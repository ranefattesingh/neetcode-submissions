func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := low + (high - low) / 2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return binarySearch(nums, low, mid - 1, target)
	}

	return binarySearch(nums, mid + 1, high, target)
}