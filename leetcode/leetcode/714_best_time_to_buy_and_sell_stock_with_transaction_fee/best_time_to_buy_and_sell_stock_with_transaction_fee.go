package leetcode714

func maxProfit(prices []int, fee int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}

	prebuy, presell := -prices[0], 0
	for i := 1; i < size; i++ {
		prebuy = max(presell-prices[i], prebuy)
		presell = max(prebuy+prices[i]-fee, presell)
	}
	return presell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
