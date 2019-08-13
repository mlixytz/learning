package leetcode338

func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>1] + (i & 1) // i的1的个数 = i/2的1的个数 + 最后一位是否为1
	}
	return dp
}
