package leetcode075

func sortColors(nums []int) {
	begin, end := 0, len(nums)-1
	for i := 0; i <= end; {
		if nums[i] == 0 {
			if i != begin {
				nums[i], nums[begin] = nums[begin], nums[i]
			}
			begin++
			i++
		} else if nums[i] == 2 {
			if i != end {
				nums[i], nums[end] = nums[end], nums[i]
			}
			end--
		} else {
			i++
		}
	}
}
