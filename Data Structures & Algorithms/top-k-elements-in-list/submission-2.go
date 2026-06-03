func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)

	for _, v := range nums {
		store[v]++
	}

	matrix := make([][]int, len(nums)+1)
	for hk, v := range store {
		matrix[v] = append(matrix[v], hk)
	}

	output := make([]int, 0, k)

	for i := len(matrix) - 1; i >= 0; i-- {
		if len(matrix[i]) > 0 {
			for j := 0; j < len(matrix[i]) && len(output) < k; j++ {
				output = append(output, matrix[i][j])
			}
		}
	}

	return output
}