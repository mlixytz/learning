package leetcode121

func maxProfit(prices []int) int {
	n := len(prices)
	m := 0
	maxlist := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			maxlist[n-i-1] = prices[n-i-1]
		} else {
			m = max(maxlist[n-i]-prices[n-i-1], m)
			maxlist[n-i-1] = max(maxlist[n-i], prices[n-i-1])
		}
	}
	return m
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
