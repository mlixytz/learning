package leetcode062

import "testing"

func TestSolution(t *testing.T) {
	if uniquePaths(3, 2) != 3 {
		t.Error("test error")
	}
	if uniquePaths(7, 3) != 28 {
		t.Error("test error")
	}
}
