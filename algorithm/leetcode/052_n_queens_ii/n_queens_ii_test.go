package leetcode052

import "testing"

func TestSolution(t *testing.T) {
	if totalNQueens(8) != 92 {
		t.Error("test error")
	}
	if totalNQueens(4) != 2 {
		t.Error("test error")
	}
}
