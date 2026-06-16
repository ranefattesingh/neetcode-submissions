func longestConsecutive(nums []int) int {
	hashMap := make(map[int]struct{})

	for  _, num := range nums {
		hashMap[num] = struct{}{}
	}

	res := 0
	for num, _ := range hashMap {
		_, ok :=  hashMap[num-1]
		if ok {
			continue
		}

		i := 1
		for _, ok := hashMap[num + i]; ok; _, ok = hashMap[num + i] {
			i++
		}

		res = max(i, res)
	}

	return res
}
