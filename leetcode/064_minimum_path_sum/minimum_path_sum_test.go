package leetcode064

import "testing"

func TestSolution(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	if minPathSum(grid) != 7 {
		t.Error("test error")
	}
}
