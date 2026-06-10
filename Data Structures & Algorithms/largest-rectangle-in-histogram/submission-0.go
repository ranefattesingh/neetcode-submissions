
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0)
	maxArea := 0

	// [7,1,7,2,2,4]

	for i := 0; i <= n; i++ {
		for len(stack) > 0 && (i == n || heights[stack[len(stack)-1]] >= heights[i]){
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			
			w := i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}

			area := h * w
			maxArea =  max(maxArea, area)
		}

		if i < n {
			stack = append(stack, i)
		}

	}

	return maxArea
}
