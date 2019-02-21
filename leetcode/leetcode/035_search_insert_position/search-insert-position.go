package leetcode035

func searchInsert(nums []int, target int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		m := l + (h-l)/2
		if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			h = m - 1
		} else {
			return m
		}
	}
	return l
}
