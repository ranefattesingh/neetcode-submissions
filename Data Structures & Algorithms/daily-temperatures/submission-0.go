func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(temperatures))
	
	for i, v := range temperatures {
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
				index := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result[index] = i - index
			}

			stack = append(stack, i)
	}
		
	return result
}


