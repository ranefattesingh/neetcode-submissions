func maxProfit(prices []int) int {
    left := 0
    profit := 0

    for right := 0; right < len(prices); right++ {
        if prices[right] < prices[left] {
            left = right
        } else {
            currentProfit := prices[right] - prices[left]
            profit = max(profit, currentProfit)
        }
    }

    return profit
}
