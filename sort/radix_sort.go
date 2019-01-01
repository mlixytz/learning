package sort

/*
   基数排序属于线性排序，时间复杂度为O(n)。（n较大，不一定优于快排）
   基于计数排序，从低位到高位排序。
*/

func radixSort(nums []int, d int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for k := 0; k < d; k++ {
		output := make([]int, n)
		tmp := make([]int, 11)

		for i := 0; i < n; i++ {
			key := getNumInPosition(nums[i], k+1)
			tmp[key]++
		}

		for i := 1; i < 11; i++ {
			tmp[i] = tmp[i] + tmp[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			key := getNumInPosition(nums[i], k+1)
			output[tmp[key]-1] = nums[i]
			tmp[key]--
		}
		nums = output
	}
	return nums
}

func getNumInPosition(num, i int) int {
	for j := 0; j < i-1; j++ {
		num /= 10
	}
	return num % 10
}
