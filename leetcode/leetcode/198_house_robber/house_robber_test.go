package leetcode198

import "testing"

func TestSolution(t *testing.T) {
	if rob([]int{1, 2, 3, 1}) != 4 {
		t.Error("test error")
	}
	if rob([]int{2, 7, 9, 3, 1}) != 12 {
		t.Error("test error")
	}
	if rob([]int{2, 1, 1, 2}) != 4 {
		t.Error("test error")
	}
}
