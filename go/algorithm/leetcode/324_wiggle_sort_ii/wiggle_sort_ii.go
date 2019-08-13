package leetcode324

func wiggleSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	m := findKthLargest(nums, (n+1)/2)
	left, right := 0, n-1
	for i := 0; i <= right; {
		nexIndexI := newIndex(i, n)
		if nums[nexIndexI] > m {
			newindexLeft := newIndex(left, n)
			nums[nexIndexI], nums[newindexLeft] = nums[newindexLeft], nums[nexIndexI]
			left++
			i++
		} else if nums[nexIndexI] < m {
			newindexRight := newIndex(right, n)
			nums[nexIndexI], nums[newindexRight] = nums[newindexRight], nums[nexIndexI]
			right--
		} else {
			i++
		}
	}
}

// 计算新索引，类似荷兰国旗问题
func newIndex(i, n int) int {
	return (1 + 2*i) % (n | 1)
}

func findKthLargest(nums []int, k int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	heap := tmp[:k]
	for i := k/2 - 1; i >= 0; i-- {
		minHeapify(heap, i)
	}
	for i := k; i < len(tmp); i++ {
		if tmp[i] > heap[0] {
			heap[0] = tmp[i]
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
