func largestRectangleArea(heights []int) int {
	stack := []int{}
	area := 0
	N := len(heights)

	for i := 0; i <= N; i++ {
		for len(stack) > 0 && (i == N || heights[stack[len(stack)-1]] >= heights[i]) {
			smallest := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area = max(area, heights[smallest]*width)
		}

		if i < N {
			stack = append(stack, i)
		}
	}

	return area
}