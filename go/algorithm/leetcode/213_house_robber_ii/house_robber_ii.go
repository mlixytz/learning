package leetcode213

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	return max(f(nums, 0, size-2), f(nums, 1, size-1))
}

func f(nums []int, a, b int) int {
	var pre1, pre2 int
	for i := a; i <= b; i++ {
		cur := max(pre1, pre2+nums[i])
		pre2, pre1 = pre1, cur
	}
	return pre1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
