package leetcode790

/*
dp[n] = dp[n-1] + dp[n-2] + 2 * (dp[n-3]+...+dp[0])
		||
dp[n] = 2*dp[n-1] + dp[n-3]
*/
const mod = 1e9 + 7

func numTilings(N int) int {
	dp1, dp2, dp3 := 1, 2, 5
	if N == 1 {
		return dp1
	}
	if N == 2 {
		return dp2
	}
	if N == 3 {
		return dp3
	}
	dp4 := 0
	for i := 4; i <= N; i++ {
		dp4 = (2*dp3 + dp1) % mod
		dp1, dp2, dp3 = dp2, dp3, dp4
	}
	return dp4
}
