package leetcode152

// dp solution
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]

	m := nums[0]
	for j := 1; j < n; j++ {
		dp[j] = nums[j]
		m = max(dp[j], m)
		for i := 0; i < j; i++ {
			dp[i] *= nums[j]
			m = max(dp[i], m)
		}
	}
	return m
}

// best solution
func maxProductBS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	r := nums[0]
	for i, imin, imax := 1, r, r; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}

		imax = max(nums[i], imax*nums[i])
		imin = min(nums[i], imin*nums[i])

		r = max(imax, r)
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
