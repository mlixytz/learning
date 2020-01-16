package leetcode018

import "sort"

func fourSum(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		for i < n-3 && i > 0 && nums[i] == nums[i-1] {
			i++
		}
		if i >= n-3 || nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			for j < n-2 && j > i+1 && nums[j] == nums[j-1] {
				j++
			}
			r := target - nums[i] - nums[j]
			s := j + 1
			e := n - 1
			for s < e {
				sum := nums[s] + nums[e]
				if sum < r {
					for s < e && nums[s+1] == nums[s] {
						s++
					}
					s++
				} else if sum > r {
					for s < e && nums[e-1] == nums[e] {
						e--
					}
					e--
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[s], nums[e]})
					for s < e && nums[s+1] == nums[s] {
						s++
					}
					s++
					for s < e && nums[e-1] == nums[e] {
						e--
					}
					e--
				}
			}
		}
	}
	return ret
}
