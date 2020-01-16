package leetcode378

import "testing"

func TestSolution(t *testing.T) {
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	if kthSmallest(matrix, 8) != 13 {
		t.Error("test error")
	}
	matrix2 := [][]int{{1, 2}, {1, 3}}
	if kthSmallest(matrix2, 3) != 2 {
		t.Error("test error")
	}
	matrix3 := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	if kthSmallest(matrix3, 8) != 13 {
		t.Error("test error")
	}
}
