func maxProfit(prices []int) int {
	l, maxProfit := 0, 0

	for r := 1; r < len(prices); r++ {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxProfit = max(maxProfit, profit)
		} else {
			l = r
		}
	}

	return maxProfit
}
