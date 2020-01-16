package leetcode746

import "testing"

func TestSolution(t *testing.T) {
	if minCostClimbingStairs([]int{10, 15, 20}) != 15 {
		t.Error("test error")
	}
	if minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}) != 6 {
		t.Error("test error")
	}
}
