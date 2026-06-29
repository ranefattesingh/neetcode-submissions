func largestRectangleArea(heights []int) int {
    stack := []int{}
    n := len(heights)
    area := 0

    for i := 0; i <= len(heights); i++ {
        for len(stack) > 0 && (i == n || heights[stack[len(stack)-1]] >= heights[i]) {
            smallest := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            width := i
            if len(stack) > 0 {
                width = (i - 1) - stack[len(stack)-1]
            }

            area = max(area, heights[smallest]*width)
        }

        if i < len(heights) {
            stack = append(stack, i)
        }
    }

    return area
}
