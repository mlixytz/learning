package leetcode376

func wiggleMaxLength(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	up, down := 1, 1

	for i := 1; i < size; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}
