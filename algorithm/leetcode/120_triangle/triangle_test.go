package leetcode120

import (
	"testing"
)

func TestSolution(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	if minimumTotal(triangle) != 11 {
		t.Error("test error")
	}
	if minimumTotalBS(triangle) != 11 {
		t.Error("test error")
	}
}
