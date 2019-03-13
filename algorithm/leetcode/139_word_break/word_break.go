package leetcode139

func wordBreak(s string, wordDict []string) bool {
	Len := len(s)
	dp := make([]bool, Len+1)
	dp[0] = true
	for i := 1; i <= Len; i++ {
		for _, val := range wordDict {
			l := len(val)
			if i >= l && s[i-l:i] == val {
				dp[i] = dp[i] || dp[i-l]
			}
		}
	}
	return dp[Len]
}
