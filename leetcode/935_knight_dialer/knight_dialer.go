package leetcode935

func knightDialer(N int) int {
	if N == 1 {
		return 10
	}
	dp := make([][]int, 10)
	for i := 0; i < 10; i++ {
		dp[i] = make([]int, N)
		dp[i][0] = 1
	}
	for k := 1; k < N; k++ {
		dp[1][k] = (dp[6][k-1] + dp[8][k-1]) % (1E9 + 7)
		dp[2][k] = (dp[7][k-1] + dp[9][k-1]) % (1E9 + 7)
		dp[3][k] = (dp[4][k-1] + dp[8][k-1]) % (1E9 + 7)
		dp[4][k] = (dp[0][k-1] + dp[3][k-1] + dp[9][k-1]) % (1E9 + 7)
		dp[6][k] = (dp[0][k-1] + dp[1][k-1] + dp[7][k-1]) % (1E9 + 7)
		dp[7][k] = (dp[2][k-1] + dp[6][k-1]) % (1E9 + 7)
		dp[8][k] = (dp[1][k-1] + dp[3][k-1]) % (1E9 + 7)
		dp[9][k] = (dp[2][k-1] + dp[4][k-1]) % (1E9 + 7)
		dp[0][k] = (dp[4][k-1] + dp[6][k-1]) % (1E9 + 7)
	}
	sum := 0
	for i := 0; i < 10; i++ {
		sum += dp[i][N-1]
	}
	return sum % (1E9 + 7)
}
