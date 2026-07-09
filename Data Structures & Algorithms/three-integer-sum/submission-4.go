func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	for i, num := range nums {
		if i > 0 && nums[i-1] == num {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := num + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{num, nums[left], nums[right]})

				currLeft := nums[left]
				for nums[left] == currLeft && left < right {
					left++
				}

				currRight := nums[right]
				for nums[right] == currRight && right > left {
					right--
				}

			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
