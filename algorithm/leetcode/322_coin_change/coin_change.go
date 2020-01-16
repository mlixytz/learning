package leetcode322

import (
	"sort"
)

func coinChange(coins []int, amount int) int {
	n := len(coins)
	if n == 0 {
		return -1
	}
	sort.Ints(coins)
	dp := make([]int, amount+1)
	dp[0] = 0
	for w := 1; w <= amount; w++ {
		dp[w] = amount + 1
		for i := 0; i < n && coins[i] <= w; i++ {
			if dp[w] > dp[w-coins[i]]+1 {
				dp[w] = dp[w-coins[i]] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
