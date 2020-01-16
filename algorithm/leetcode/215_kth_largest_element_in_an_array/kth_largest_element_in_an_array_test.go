package leetcode215

import "testing"

func TestSolution(t *testing.T) {
	if findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) != 5 {
		t.Error("test error")
	}
	if findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) != 4 {
		t.Error("test error")
	}
}
