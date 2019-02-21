package leetcode532

import "testing"

func TestSolution(t *testing.T) {
	if !checkSubarraySum([]int{23, 2, 4, 6, 7}, 6) {
		t.Error("test error")
	}
	if !checkSubarraySum([]int{23, 2, 6, 4, 7}, 6) {
		t.Error("test error")
	}
	if !checkSubarraySum([]int{0, 0}, 0) {
		t.Error("test error")
	}
}
