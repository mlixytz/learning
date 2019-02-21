package leetcode516

func longestPalindromeSubseq(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for Len := 2; Len <= n; Len++ {
		for left := 0; left+Len-1 < n; left++ {
			right := left + Len - 1
			if s[left] == s[right] {
				dp[left][right] = dp[left+1][right-1] + 2
			} else {
				dp[left][right] = max(dp[left+1][right], dp[left][right-1])
			}
		}
	}
	return dp[0][n-1]
}
