package leetcode031

import "sort"

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
				sort.Ints(nums[i+1 : n])
				return
			}
		}
	}
	sort.Ints(nums)
}
