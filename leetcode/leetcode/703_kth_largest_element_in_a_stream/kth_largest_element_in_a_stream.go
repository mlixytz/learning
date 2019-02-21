package leetcode703

type KthLargest struct {
	Heap []int
	Size int
}

func Constructor(k int, nums []int) KthLargest {
	var heap []int
	if len(nums) < k {
		heap = nums
	} else {
		heap = nums[:k]
	}
	for i := len(heap)/2 - 1; i >= 0; i-- {
		minHeapify(heap, i)
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			minHeapify(heap, 0)
		}
	}
	return KthLargest{heap, k}
}

func (this *KthLargest) Add(val int) int {
	if len(this.Heap) < this.Size {
		this.Heap = append(this.Heap, val)
		index := len(this.Heap) - 1
		for index > 0 && this.Heap[parent(index)] > this.Heap[index] {
			this.Heap[parent(index)], this.Heap[index] = this.Heap[index], this.Heap[parent(index)]
			index = parent(index)
		}
	} else {
		if val > this.Heap[0] {
			this.Heap[0] = val
			minHeapify(this.Heap, 0)
		}
	}
	return this.Heap[0]
}

func parent(i int) int {
	if i%2 == 0 {
		return (i - 2) / 2
	}
	return (i - 1) / 2
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

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
