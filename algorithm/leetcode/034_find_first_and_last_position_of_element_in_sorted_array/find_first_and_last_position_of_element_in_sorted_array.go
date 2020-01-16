package leetcode034

// best solution
func searchRangeBestSolution(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, h := 0, len(nums)-1
	// find first one
	for l < h {
		m := l + (h-l)/2
		if nums[m] < target {
			l = m + 1
		} else {
			h = m
		}
	}
	if nums[l] != target {
		return []int{-1, -1}
	}
	start := l

	// find last one
	h = len(nums) - 1
	for l < h {
		m := l + (h-l+1)/2
		if nums[m] > target {
			h = m - 1
		} else {
			l = m
		}
	}

	return []int{start, l}
}

// my solution
func searchRange(nums []int, target int) []int {
	l, h := 0, len(nums)-1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] > target {
			h = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			s, e, changed := m, m, true
			for changed {
				changed = false
				if s > 0 && nums[s-1] == target {
					s--
					changed = true
				}

				if e < len(nums)-1 && nums[e+1] == target {
					e++
					changed = true
				}
			}
			return []int{s, e}
		}
	}
	return []int{-1, -1}
}
