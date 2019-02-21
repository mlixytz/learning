package leetcode416

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
		if dp[sum] {
			return true
		}
	}
	return dp[sum]
}
