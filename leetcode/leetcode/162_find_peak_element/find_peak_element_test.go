package leetcode162

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if findPeakElement([]int{1, 2, 3, 1}) != 2 {
		t.Error("test error")
	}
	if findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}) != 1 && findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}) != 5 {
		t.Error("test error")
	}
}
