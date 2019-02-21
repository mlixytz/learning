package leetcode153

func findMin(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2
		if nums[l] <= nums[m] && nums[h] <= nums[m] {
			l = m + 1
		} else {
			h = m
		}
	}
	return nums[h]
}
