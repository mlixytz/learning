package leetcode053

func maxSubArray(nums []int) int {
	m, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		m = max(m, sum)
	}
	return m
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
