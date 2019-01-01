package sort

/*
   桶排序属于线性排序，时间复杂度为O(n)
   需满足：所有桶的大小的平方和与总的元素数呈线性关系
*/

func bucketSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	// 找到数组最大值和最小值
	minValue, maxValue := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if minValue > nums[i] {
			minValue = nums[i]
		} else if maxValue < nums[i] {
			maxValue = nums[i]
		}
	}

	// 桶初始化
	bucketSize := 5
	bucketCount := (maxValue-minValue)/bucketSize + 1
	buckets := make([][]int, bucketCount)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = make([]int, 0)
	}

	// 利用映射函数将数据映射到各个桶中
	for i := 0; i < n; i++ {
		key := (nums[i] - minValue) / bucketSize
		buckets[key] = append(buckets[key], nums[i])
	}

	output := make([]int, 0)
	// 桶内排序
	for i := 0; i < len(buckets); i++ {
		l := len(buckets[i])
		if l > 1 {
			insertionSort(buckets[i])
		}
		for j := 0; j < len(buckets[i]); j++ {
			output = append(output, buckets[i][j])
		}
	}
	return output
}
