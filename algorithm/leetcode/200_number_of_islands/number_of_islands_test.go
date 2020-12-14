package leetcode200

import "testing"

func TestSolution(t *testing.T) {
	grid := [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}
	if numIslands(grid) != 1 {
		t.Error("test error")
	}
}
