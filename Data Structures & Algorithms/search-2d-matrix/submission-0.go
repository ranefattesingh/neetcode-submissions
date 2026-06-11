func searchMatrix(matrix [][]int, target int) bool {
	t, b := 0, len(matrix) - 1

	for t <= b {
		vm := t + (b-t)/2
		if matrix[vm][0] > target {
			b = vm - 1
		} else if matrix[vm][len(matrix[vm])-1] < target {
			t = vm + 1
		} else {
			l, r := 0, len(matrix[vm]) - 1
			for l <= r {
				hm := l + (r-l)/2
				if matrix[vm][hm] > target {
					r = hm - 1
				} else if matrix[vm][hm] < target {
					l = hm + 1
				} else {
					return true
				}
			}

			return false
		}
	}

	return false
}
