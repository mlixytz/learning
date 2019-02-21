package leetcode074

import "testing"

func TestSolution(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	if searchMatrix(matrix, 3) != true {
		t.Error("test error")
	}
	if searchMatrix(matrix, 13) != false {
		t.Error("test error")
	}
	if searchMatrix(matrix, 0) != false {
		t.Error("test error")
	}
	if searchMatrix(matrix, 60) != false {
		t.Error("test error")
	}
}
