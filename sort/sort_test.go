package sort

import (
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := randArray(10)
	t.Log("initial array is", nums)
	nums = bubbleSort(nums)
	t.Log("sorted array is", nums)
	if !isSortedArray(nums) {
		t.Error("bubble sort test error")
	}
}

func TestSelectionSort(t *testing.T) {
	nums := randArray(10)
	t.Log("initial array is", nums)
	nums = selectionSort(nums)
	t.Log("sorted array is", nums)
	if !isSortedArray(nums) {
		t.Error("selection sort test error")
	}
}

func TestInsertionSort(t *testing.T) {
	nums := randArray(10)
	t.Log("initial array is", nums)
	nums = insertionSort(nums)
	t.Log("sorted array is", nums)
	if !isSortedArray(nums) {
		t.Error("insertion sort test error")
	}
}

func TestShellSort(t *testing.T) {
	nums := randArray(10)
	t.Log("initial array is", nums)
	nums = shellSort(nums)
	t.Log("sorted array is", nums)
	if !isSortedArray(nums) {
		t.Error("shell sort test error")
	}
}

func TestMergeSort(t *testing.T) {
	nums := randArray(10)
	t.Log("initial array is", nums)
	nums = mergeSort(nums)
	t.Log("sorted array is", nums)
	if !isSortedArray(nums) {
		t.Error("merge sort test error")
	}
}

func TestQuickSort(t *testing.T) {
	nums := randArray(10)
	t.Log("initial array is", nums)
	nums = quickSort(nums)
	t.Log("sorted array is", nums)
	if !isSortedArray(nums) {
		t.Error("quick sort test error")
	}
}

func isSortedArray(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func randArray(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(1000)
	}
	return nums
}
