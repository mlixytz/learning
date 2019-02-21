package leetcode153

import "testing"

func TestSolution(t *testing.T) {
	if findMin([]int{3, 4, 5, 1, 2}) != 1 {
		t.Error("test error")
	}
	if findMin([]int{4, 5, 6, 7, 0, 1, 2}) != 0 {
		t.Error("test error")
	}
	if findMin([]int{2, 1}) != 1 {
		t.Error("test error")
	}
}
