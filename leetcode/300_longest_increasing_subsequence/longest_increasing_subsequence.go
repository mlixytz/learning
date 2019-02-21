package leetcode300

// dp solution
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	_max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxLen := 1
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = _max(dp[j]+1, dp[i])
			}
		}
		maxLen = _max(dp[i], maxLen)
	}
	return maxLen
}

// dp with binary search solution
func lengthOfLISBS(nums []int) int {
	dp := make([]int, len(nums))
	size := 0
	for _, val := range nums {
		i, j := 0, size
		for i < j {
			m := i + (j-i)/2
			if dp[m] < val {
				i = m + 1
			} else {
				j = m
			}
		}
		dp[i] = val
		if i == size {
			size++
		}
	}
	return size
}
