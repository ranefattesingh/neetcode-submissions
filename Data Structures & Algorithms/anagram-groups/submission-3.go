func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)

	for _, v  := range strs  {
		k := [26]int{}
		for _, m := range v {
			k[m-'a']++
		}

		_, ok := hashMap[k]
		if !ok {
			hashMap[k] = make([]string, 0)
		}
		
		hashMap[k] = append(hashMap[k],  v)
	}

	matrix := make([][]string, 0, len(hashMap))
	for _, v := range hashMap {
		matrix = append(matrix, v)
	}

	return matrix
}