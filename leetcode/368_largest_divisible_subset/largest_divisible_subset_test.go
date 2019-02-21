package leetcode368

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if !compareSlice(largestDivisibleSubset([]int{1, 2, 3}), []int{1, 2}) && !compareSlice(largestDivisibleSubset([]int{1, 2, 3}), []int{1, 3}) &&
		!compareSlice(largestDivisibleSubset([]int{1, 2, 3}), []int{2, 1}) && !compareSlice(largestDivisibleSubset([]int{1, 2, 3}), []int{3, 1}) {
		t.Error("test error")
	}
	if !compareSlice(largestDivisibleSubset([]int{1, 2, 4, 8}), []int{1, 2, 4, 8}) && !compareSlice(largestDivisibleSubset([]int{1, 2, 4, 8}), []int{8, 4, 2, 1}) {
		t.Error("test error")
	}
	if !compareSlice(largestDivisibleSubset([]int{3, 4, 16, 8}), []int{16, 8, 4}) {
		t.Error("test error")
	}
}

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
