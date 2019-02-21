package leetcode852

import "testing"

func TestSolution(t *testing.T) {
	if peakIndexInMountainArray([]int{0, 1, 0}) != 1 {
		t.Error("test error")
	}
	if peakIndexInMountainArray([]int{0, 2, 1, 0}) != 1 {
		t.Error("test error")
	}
}
