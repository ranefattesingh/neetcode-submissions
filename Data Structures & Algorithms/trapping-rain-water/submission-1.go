func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax := height[l]
	rightMax := height[r]
	water := 0

	for l < r {
		if leftMax <= rightMax {
			l++
			leftMax = max(leftMax, height[l])
			water += (leftMax - height[l])
		} else {
			r--
			rightMax = max(rightMax, height[r])
			water += (rightMax - height[r])
		}
	}

	return water
}