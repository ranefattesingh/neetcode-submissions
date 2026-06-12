func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, v := range piles {
		r = max(r, v)
	} 

	globalMin := r

	for l <= r {
		m := l + (r-l)/2

		sum := 0
		for i := 0; i < len(piles); i++ {
			sum += int(math.Ceil(float64(piles[i])/float64(m)))
		}

		if sum <= h {
			r = m - 1
			globalMin = m
		} else {
			l = m + 1
		}
	}

	return globalMin
}
