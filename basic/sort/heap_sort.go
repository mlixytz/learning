package sort

func heapSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	i := n
	for i > 1 {
		buildMaxHeap(nums[:i])
		nums[0], nums[i-1] = nums[i-1], nums[0]
		i--
	}
	return nums
}

func buildMaxHeap(nums []int) {
	n := len(nums)
	// n/2 + 1 ... n 全部属于叶节点
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i)
	}
}

func maxHeapify(nums []int, i int) {
	n := len(nums)
	left := i << 1
	largest := i
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	right := left + 1
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest)
	}
}
