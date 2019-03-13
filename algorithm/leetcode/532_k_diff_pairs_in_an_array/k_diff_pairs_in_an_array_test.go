package leetcode532

import "testing"

func TestSolution(t *testing.T) {
	if findPairs([]int{3, 1, 4, 1, 5}, 2) != 2 {
		t.Error("test error")
	}
	if findPairs([]int{1, 2, 3, 4, 5}, 1) != 4 {
		t.Error("test error")
	}
	if findPairs([]int{1, 3, 1, 5, 4}, 0) != 1 {
		t.Error("test error")
	}
}
