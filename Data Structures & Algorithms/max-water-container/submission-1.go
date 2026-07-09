func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1
	maxArea := 0

	for left < right {
		height := min(heights[left], heights[right])
		area := height * (right - left)
		maxArea = max(maxArea, area)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
