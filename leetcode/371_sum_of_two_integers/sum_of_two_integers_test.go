package leetcode371

import "testing"

func TestSolution(t *testing.T) {
	if getSum(7, 3) != 10 {
		t.Error("test error")
	}
	if getSum(1, 2) != 3 {
		t.Error("test error")
	}
	if getSum(-2, 3) != 1 {
		t.Error("test error")
	}
}
