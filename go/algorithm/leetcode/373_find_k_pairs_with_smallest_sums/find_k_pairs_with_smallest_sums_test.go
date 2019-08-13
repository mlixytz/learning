package leetcode373

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if !isEqual(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3), [][]int{{1, 2}, {1, 4}, {1, 6}}) {
		t.Error("test error")
	}
	if !isEqual(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2), [][]int{{1, 1}, {1, 1}}) {
		t.Error("test error")
	}
	if !isEqual(kSmallestPairs([]int{}, []int{}, 5), [][]int{}) {
		t.Error("test error")
	}
	if !isEqual(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 10), [][]int{{1, 1}, {1, 1}, {1, 2}, {1, 2}, {2, 1}, {1, 3}, {1, 3}, {2, 2}, {2, 3}}) {
		t.Error("test error")
	}
}

func isEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
