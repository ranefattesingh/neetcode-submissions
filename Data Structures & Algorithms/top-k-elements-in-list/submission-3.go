func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	matrix :=  make([][]int, len(nums)+1)

	for  _, v := range nums {
		hashMap[v]++
	}


	for num, count := range hashMap {
		matrix[count] = append(matrix[count], num)
	}

	res := []int{}
	for i := len(matrix) - 1; i > 0; i-- {
		for _, num := range matrix[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

/*

	INPUT =  [4, 5, 5, 6, 6, 6, 8, 8], k = 2
	STEP 1: CREATE COUNT HASH MAP => (hashMap[num] -> count)
	[4] = 1
	[5] = 2
	[6] = 3

	STEP 2: CREATE FREQUENCY MATRIX => (matrix[count] -> [num1, num2])
	[0] -> nil
	[1] -> [4]
	[2] -> [5, 8]
	[3] -> [6]

	STEP 3: SELECT FROM TOP INDICES TOP k ELEMENTS
	[3] ->  [6]
	[2]  -> [5]

	return [5, 6]
*/