func groupAnagrams(strs []string) [][]string {
	store := make(map[[26]int][]string)
	for _, v := range strs {
		count := [26]int{}
		for _, c := range v {
			index := int(c) - int('a')
			count[index]++
		}

		store[count] = append(store[count], v)
	}

	matrix := make([][]string, 0, len(store))
	for _, v := range store {
		matrix = append(matrix, v)
	}

	return matrix
}