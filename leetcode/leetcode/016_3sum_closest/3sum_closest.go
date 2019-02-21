package leetcode016

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := 1 << 31
	for k, val := range nums {
		for i, j := 0, len(nums)-1; i < j; {
			if i == k {
				i++
			} else if j == k {
				j--
			} else if nums[i]+nums[j]+val > target {
				if abs(ans-target) > abs(nums[i]+nums[j]+val-target) {
					ans = nums[i] + nums[j] + val
				}
				j--
			} else if nums[i]+nums[j]+val < target {
				if abs(ans-target) > abs(nums[i]+nums[j]+val-target) {
					ans = nums[i] + nums[j] + val
				}
				i++
			} else {
				return target
			}
		}
	}
	return ans
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
