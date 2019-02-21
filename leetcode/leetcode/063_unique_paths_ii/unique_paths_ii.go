package leetcode063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[j] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ {
		if dp[0] == 0 || obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[j] = dp[j-1] + dp[j]
			} else {
				dp[j] = 0
			}
		}
	}
	return dp[n-1]
}
