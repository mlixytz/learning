package leetcode769

import "testing"

func TestSolution(t *testing.T) {
	if maxChunksToSorted([]int{1, 2, 0, 3}) != 2 {
		t.Error("test error")
	}
}
