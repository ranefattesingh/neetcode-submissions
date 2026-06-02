func getConcatenation(nums []int) []int {
    n := len(nums)
	ans := make([]int, 0, n*2)
	i := 0

	for j := 0; j < n*2; j++ {
		ans = append(ans, nums[i%n])
		i++
	}

	return ans
}
