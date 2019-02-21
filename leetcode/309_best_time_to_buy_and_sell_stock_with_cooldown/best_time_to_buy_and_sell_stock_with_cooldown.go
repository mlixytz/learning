package leetcode309

func maxProfit(prices []int) int {
	_max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	size := len(prices)
	if size < 2 {
		return 0
	}
	prevBuy, prevSell, sell := -prices[0], 0, 0
	for i := 1; i < size; i++ {
		tmp := sell
		sell = _max(prevBuy+prices[i], sell)
		prevBuy = _max(prevSell-prices[i], prevBuy)
		prevSell = tmp
	}
	return sell
}
