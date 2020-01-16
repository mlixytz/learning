package leetcode064

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	sums := make([][]int, m)

	for i := 0; i < m; i++ {
		sums[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				sums[i][j] = min(sums[i-1][j], sums[i][j-1]) + grid[i][j]
			} else if i > 0 && j == 0 {
				sums[i][j] = sums[i-1][j] + grid[i][j]
			} else if i == 0 && j > 0 {
				sums[i][j] = sums[i][j-1] + grid[i][j]
			} else {
				sums[i][j] = grid[i][j]
			}
		}
	}
	return sums[m-1][n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
