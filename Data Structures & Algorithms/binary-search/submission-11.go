func search(nums []int, target int) int {
	var binarySearch func(nums []int, target, low, high int) int
	
	binarySearch = func(nums []int, target, low, high int) int {
		if low > high {
			return -1
		}

		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			return binarySearch(nums, target, low, mid-1)
		}

		return binarySearch(nums, target, mid+1, high)
	}

	return binarySearch(nums, target, 0, len(nums)-1)
}

