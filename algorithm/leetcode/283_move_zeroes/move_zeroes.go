package leetcode283

func moveZeroes(nums []int) {
	i := 0 // zero index
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
