package leetcode1051

import "testing"

func TestSolution(t *testing.T) {
	if heightChecker([]int{1, 1, 4, 2, 1, 3}) != 3 {
		t.Error("test error")
	}
}
