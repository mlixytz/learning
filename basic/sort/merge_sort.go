package sort

func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	return merge(mergeSort(nums[:n/2]), mergeSort(nums[n/2:]))
}

func merge(left, right []int) []int {
	l := len(left) + len(right)
	result := make([]int, 0, l)
	var i, j int
	for k := 0; k < l; k++ {
		if i == len(left) {
			return append(result, right[j:]...)
		}
		if j == len(right) {
			return append(result, left[i:]...)
		}
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return result
}
