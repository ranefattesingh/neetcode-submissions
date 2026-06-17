func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

    for l < r {
        m := (l + r) / 2
        if nums[m] > nums[r] {
            l = m + 1
        } else {
            r = m
        }
    }

    pivot := l

	var binarySearch func(left, right int) int
    binarySearch = func(left, right int) int {
        for left <= right {
            mid := (left + right) / 2
            if nums[mid] == target {
                return mid
            } else if nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        return -1
    }

    result := binarySearch(0, pivot-1)
    if result != -1 {
        return result
    }

    return binarySearch(pivot, len(nums)-1)
}
