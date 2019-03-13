package leetcode033

func search(nums []int, target int) int {
	maxIndex := len(nums) - 1
	if maxIndex < 0 {
		return -1
	}
	l, h := 0, maxIndex
	for l < h {
		m := l + (h-l)/2
		if nums[m] > nums[h] {
			l = m + 1
		} else {
			h = m
		}
	}
	if target > nums[maxIndex] {
		l, h = 0, l-1
	} else if target < nums[maxIndex] {
		h = maxIndex
	} else {
		return maxIndex
	}

	for l <= h {
		m := l + (h-l)/2
		if nums[m] > target {
			h = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}
