package leetcode740

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	size := 10001
	dp := make([]int, size)
	for _, v := range nums {
		dp[v] += v
	}
	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-2]+dp[i], dp[i-1])
	}
	return dp[size-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
