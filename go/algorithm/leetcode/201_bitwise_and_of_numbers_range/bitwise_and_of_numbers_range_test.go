package leetcode201

import "testing"

func TestSolution(t *testing.T) {
	if rangeBitwiseAnd(5, 7) != 4 {
		t.Error("test error")
	}
	if rangeBitwiseAnd(0, 1) != 0 {
		t.Error("test error")
	}
}
