package leetcode198

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	} else if size == 1 {
		return nums[0]
	}
	sum := make([]int, size)
	sum[0] = nums[0]
	sum[1] = max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		sum[i] = max(sum[i-2]+nums[i], sum[i-1])
	}
	return sum[size-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
