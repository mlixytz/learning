package sort

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, a, b int) {
	if a >= b {
		return
	}
	pivotIndex := partition(nums, a, b)
	sort(nums, a, pivotIndex-1)
	sort(nums, pivotIndex+1, b)
}

func partition(nums []int, p, r int) int {
	pivot := nums[r]
	i := p - 1
	for j := p; j < r; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[r], nums[i+1] = nums[i+1], nums[r]
	return i + 1
}
