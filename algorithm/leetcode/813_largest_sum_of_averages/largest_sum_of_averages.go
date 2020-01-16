package leetcode813

func largestSumOfAverages(A []int, K int) float64 {
	size := len(A)
	sum := make([]int, size+1)
	dp := make([][]float64, size+1)
	dp[0] = make([]float64, K+1)
	for i := 1; i <= size; i++ {
		dp[i] = make([]float64, K+1)
		sum[i] = sum[i-1] + A[i-1]
		dp[i][1] = float64(sum[i]) / float64(i)
	}

	for k := 2; k <= K; k++ {
		for i := 1; i <= size; i++ {
			m := 0.0
			for j := 0; j < i; j++ {
				m = max(dp[j][k-1]+float64(sum[i]-sum[j])/float64(i-j), m)
			}
			dp[i][k] = m
		}
	}
	return dp[size][K]
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
