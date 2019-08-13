package leetcode215

func findKthLargest(nums []int, k int) int {
	heap := nums[:k]
	for i := k/2 - 1; i >= 0; i-- {
		minHeapify(heap, i)
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			minHeapify(heap, 0)
		}
	}
	return heap[0]
}

func minHeapify(nums []int, i int) {
	n := len(nums)
	left := 2*i + 1
	right := 2*i + 2
	smallest := i

	if left < n && nums[left] < nums[smallest] {
		smallest = left
	}
	if right < n && nums[right] < nums[smallest] {
		smallest = right
	}
	if smallest != i {
		nums[smallest], nums[i] = nums[i], nums[smallest]
		minHeapify(nums, smallest)
	}
}
