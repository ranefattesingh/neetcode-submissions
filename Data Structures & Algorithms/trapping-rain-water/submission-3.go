func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]

	trappedWater := 0
	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			trappedWater += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			trappedWater += rightMax - height[right]
		}
	}

	return trappedWater
}
