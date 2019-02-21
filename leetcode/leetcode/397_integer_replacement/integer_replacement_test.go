package leetcode397

import "testing"

func TestSolution(t *testing.T) {
	if integerReplacement(8) != 3 {
		t.Error("test error")
	}
	if integerReplacement(7) != 4 {
		t.Error("test error")
	}
}
