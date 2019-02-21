package leetcode016

import "testing"

func TestSolution(t *testing.T) {
	if threeSumClosest([]int{-1, 2, 1, -4}, 1) != 2 {
		t.Error("test error")
	}
}
