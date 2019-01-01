package sort

/*
   计数排序属于线性、非原址、稳定排序，时间复杂度为O(n)。
   需要知道输入数组的最大值{maxValue}和额外的空间。
   额外数组：1.输出数组，长度与输入数组相同；2.临时数组，长度为maxValue。
*/

func countingSort(nums []int, maxValue int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	output := make([]int, n)
	tmp := make([]int, maxValue+1)

	for i := 0; i < n; i++ {
		tmp[nums[i]]++
	}

	for i := 1; i <= maxValue; i++ {
		tmp[i] = tmp[i] + tmp[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[tmp[nums[i]]-1] = nums[i]
		tmp[nums[i]]--
	}
	return output
}
