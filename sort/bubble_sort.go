package sort

func bubbleSort(nums []int) []int {
	var flag bool
	for i := 0; i < len(nums); i++ {
		flag = false
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}
