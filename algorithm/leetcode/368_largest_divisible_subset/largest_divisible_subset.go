package leetcode368

import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)
	count := make([]int, len(nums))
	prev := make([]int, len(nums))
	max, index := 0, -1
	for i := 0; i < len(nums); i++ {
		count[i] = 1
		prev[i] = -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && count[j]+1 > count[i] {
				count[i] = count[j] + 1
				prev[i] = j
			}
		}
		if count[i] > max {
			max = count[i]
			index = i
		}
	}
	res := make([]int, 0)
	for index != -1 {
		res = append(res, nums[index])
		index = prev[index]
	}
	return res
}
