package leetcode646

import "testing"

func TestSolution(t *testing.T) {
	if findLongestChain([][]int{{2, 3}, {1, 2}, {3, 4}}) != 2 {
		t.Error("test error")
	}
}
