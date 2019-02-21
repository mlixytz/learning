package leetcode688

func knightProbability(N int, K int, r int, c int) float64 {
	moves := [][]int{{-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}
	dp := [25][25]float64{}
	dp[r][c] = 1.0
	for k := 0; k < K; k++ {
		dp2 := [25][25]float64{}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if dp[i][j] == 0 {
					continue
				}
				for _, val := range moves {
					row := i + val[0]
					col := j + val[1]
					if row >= 0 && row < N && col >= 0 && col < N {
						dp2[row][col] += dp[i][j] / 8.0
					}
				}
			}
		}
		dp = dp2
	}

	rate := 0.0
	for _, val := range dp {
		for _, v := range val {
			rate += v
		}
	}
	return rate
}
