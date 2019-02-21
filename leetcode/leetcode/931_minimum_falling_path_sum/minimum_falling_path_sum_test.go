package leetcode931

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}) != 12 {
		t.Error("test error")
	}
	if minFallingPathSum([][]int{{69}}) != 69 {
		t.Error("test error")
	}
}
