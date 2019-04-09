package leetcode041

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
