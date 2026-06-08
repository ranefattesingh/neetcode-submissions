func threeSum(nums []int) [][]int {
	triplets := make([][]int, 0)

	sort.Ints(nums)
	
	for i := 0; i < len(nums); i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i + 1, len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				triplets = append(triplets, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				// skip duplicates
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return triplets
}
