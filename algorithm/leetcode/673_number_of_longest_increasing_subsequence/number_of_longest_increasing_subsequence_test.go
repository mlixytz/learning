package leetcode673

import "testing"

func TestSolution(t *testing.T) {
	if findNumberOfLIS([]int{1, 3, 5, 4, 7}) != 2 {
		t.Error("test error")
	}
	if findNumberOfLIS([]int{2, 2, 2, 2, 2}) != 5 {
		t.Error("test error")
	}
	if findNumberOfLIS([]int{1, 3, 5, 3, 3}) != 1 {
		t.Error("test error")
	}
	if findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}) != 3 {
		t.Error("test error")
	}
}
