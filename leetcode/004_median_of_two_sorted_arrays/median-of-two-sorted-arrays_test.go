package leetcode004

import "testing"

func TestSolution(t *testing.T) {
	result := findMedianSortedArrays([]int{1, 2, 3}, []int{2, 3, 4})
	if result != 2.5 {
		t.Error("test error")
	}
}
