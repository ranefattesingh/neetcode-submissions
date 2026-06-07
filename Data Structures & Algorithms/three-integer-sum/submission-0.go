func threeSum(nums []int) [][]int {
	output := make([][]int, 0)
	sort.Ints(nums)

	if nums[0] > 0 {
		return [][]int{}
	}

	for index, num := range nums{
		if index > 0 && num == nums[index-1] {
			continue
		}

		l, r := index + 1, len(nums) - 1

		for l < r {
			s := num + nums[l] + nums[r]
			if s > 0 {
				r--
			} else if s < 0 {
				l++
			} else {
				output = append(output, []int{num, nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return output
}
