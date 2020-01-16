package leetcode055

func canJump(nums []int) bool {
	lastIndex := len(nums) - 1
	if lastIndex < 1 {
		return true
	}
	for i := lastIndex - 1; i >= 0; i-- {
		off := lastIndex - i
		if nums[i] >= off {
			return canJump(nums[:i+1])
		}
	}
	return false
}
