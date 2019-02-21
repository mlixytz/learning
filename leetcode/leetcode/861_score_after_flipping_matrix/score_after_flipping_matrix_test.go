package leetcode861

import "testing"

func TestSolution(t *testing.T) {
	if matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}) != 39 {
		t.Error("test error")
	}
}
