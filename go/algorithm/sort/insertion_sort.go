package sort

func insertionSort(nums []int) []int {
	var curr int
	for i := 1; i < len(nums); i++ {
		curr = nums[i]
		for j := i - 1; j >= 0; j-- {
			if curr < nums[j] {
				nums[j+1] = nums[j]
				if j == 0 {
					nums[0] = curr
				}
			} else {
				nums[j+1] = curr
				break
			}
		}
	}
	return nums
}
