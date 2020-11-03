package leetcode435

import "testing"

func TestSolution(t *testing.T) {
	if eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}) != 2 {
		t.Error("test error")
	}
}
