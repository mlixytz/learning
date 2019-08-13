package leetcode303

import "testing"

func TestSolution(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	na := Constructor(nums)
	if na.SumRange(0, 2) != 1 {
		t.Error("test error")
	}
	if na.SumRange(2, 5) != -1 {
		t.Error("test error")
	}
	if na.SumRange(0, 5) != -3 {
		t.Error("test error")
	}
}
