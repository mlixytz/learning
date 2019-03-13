package leetcode213

import "testing"

func TestSolution(t *testing.T) {
	if rob([]int{2, 3, 2}) != 3 {
		t.Error("test error")
	}
	if rob([]int{1, 2, 3, 1}) != 4 {
		t.Error("test error")
	}
}
