func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(temperatures))

	for day, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = day - top
		}

		stack = append(stack, day)
	}

	return result
}


