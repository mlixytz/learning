package leetcode221

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	sz := 0
	dp := make([]int, n)
	for i := range matrix[0] {
		if matrix[0][i] == '1' {
			dp[i] = 1
			sz = 1
		}
	}
	for i := 1; i < m; i++ {
		var mark int
		if matrix[i][0] == '1' {
			mark = 1
			sz = max(sz, 1)
		}
		for j := 1; j < n; j++ {
			tmp := dp[j]
			if matrix[i][j] == '1' {
				dp[j] = min(dp[j-1], dp[j], mark) + 1
			} else {
				dp[j] = 0
			}
			mark = tmp
			sz = max(dp[j], sz)
		}
		if matrix[i][0] == '1' {
			dp[0] = 1
		} else {
			dp[0] = 0
		}
	}
	return sz * sz
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
