package leetcode053

import "testing"

func TestSolution(t *testing.T) {
	if maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Error("test error")
	}
	if maxSubArray([]int{-2, -3, -1}) != -1 {
		t.Error("test error")
	}
}
