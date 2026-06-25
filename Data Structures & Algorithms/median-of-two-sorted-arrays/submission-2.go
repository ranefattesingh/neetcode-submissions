func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a, b := nums1, nums2
	if len(a) > len(b) {
		a, b = b, a
	}

	n, m := len(a), len(b)
	total := n + m
	half := (total + 1) / 2
	l, r := 0, n

	for l <= r {
		i := l + (r - l) / 2 // no of elements taken from 'a'
		j := half - i // no of elements taken from 'b'

		var aLeft, aRight, bLeft, bRight int

		if i > 0 {
			aLeft = a[i-1]
		} else {
			aLeft = math.MinInt64
		}

		if i < len(a) {
			aRight = a[i]
		} else {
			aRight = math.MaxInt64
		}

		if j > 0 {
			bLeft = b[j-1]
		} else {
			bLeft = math.MinInt64
		}

		if j < len(b) {
			bRight = b[j]
		} else {
			bRight = math.MaxInt64
		}

		if aLeft <= bRight && bLeft <= aRight {
			// isOdd ?
			if total % 2 != 0 {
				return float64(max(aLeft, bLeft))
			}

			return float64(max(aLeft, bLeft) + min(aRight, bRight)) / 2.0
		} else if aLeft > bRight {
			// reduce no of elements from A
			r = i - 1
		} else {
			l = i + 1
		}
	}

	return -1
}