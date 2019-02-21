package leetcode474

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, val := range strs {
		_m, _n := 0, 0
		for _, _val := range val {
			if _val == '1' {
				_n++
			} else {
				_m++
			}
		}
		for i := m; i >= _m; i-- {
			for j := n; j >= _n; j-- {
				dp[i][j] = max(dp[i][j], dp[i-_m][j-_n]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
