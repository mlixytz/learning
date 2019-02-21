package leetcode081

import (
	"fmt"
)

// best solution
func searchBestSolution(nums []int, target int) bool {
	lenNums := len(nums)
	if lenNums == 0 {
		return false
	}
	l, h := 0, lenNums-1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] == target {
			return true
		}
		if nums[l] == nums[m] && nums[h] == nums[m] {
			l++
			h--
		} else if nums[l] <= nums[m] {
			if nums[l] <= target && nums[m] > target {
				h = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[h] >= target && nums[m] < target {
				l = m + 1
			} else {
				h = m - 1
			}
		}
	}
	return false
}

// my solution
func search(nums []int, target int) bool {
	maxIndex := len(nums) - 1
	if maxIndex < 0 {
		return false
	}
	l, h := 0, maxIndex
	for l < h {
		m := l + (h-l)/2
		fmt.Println(l, h, m)
		if nums[m] > nums[h] {
			l = m + 1
		} else if nums[m] < nums[h] {
			h = m
		} else {
			flag, temph := false, h
			for i := m + 1; i <= temph; i++ {
				if nums[m] < nums[i] {
					l = m + 1
					break
				} else if nums[m] > nums[i] {
					l = i
					flag = true
					break
				} else {
					h = m
				}
			}
			if flag {
				break
			}
		}
	}

	if target > nums[maxIndex] {
		l, h = 0, l-1
	} else if target < nums[maxIndex] {
		h = maxIndex
	} else {
		return true
	}
	for l <= h {
		m := l + (h-l)/2
		if nums[m] > target {
			h = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return true
		}
	}
	return false
}
