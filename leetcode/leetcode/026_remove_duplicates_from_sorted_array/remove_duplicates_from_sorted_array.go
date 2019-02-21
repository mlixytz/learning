package leetcode026

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	j, tmp := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != tmp {
			tmp = nums[i]
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return j
}
