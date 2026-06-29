func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	var (
		total = len(nums1) + len(nums2)
		low = 0
		high = len(nums1)
		half = (total + 1) / 2
	)

	for low <= high {
		i := (high + low)/2
		j := half - i

		l1 := math.MinInt
		if i > 0 {
			l1 = nums1[i-1]
		}

		r1 := math.MaxInt
		if i < len(nums1) {
			r1 = nums1[i]
		}

		l2 := math.MinInt
		if j > 0 {
			l2 = nums2[j-1]
		}

		r2 := math.MaxInt
		if j < len(nums2) {
			r2 = nums2[j]
		}

		if l1 <= r2 && l2 <= r1 {
			if total%2 == 0 {
				return float64(max(l1, l2)+min(r1, r2)) / float64(2)
			}

			return float64(max(l1, l2))
		} else if l1 > r2 {
			high = i - 1
		} else {
			low = i + 1
		}
	}

	return -1
}
