package leetcode476

import "testing"

func TestSolution(t *testing.T) {
	if findComplement(5) != 2 {
		t.Error("test error")
	}
	if findComplement(1) != 0 {
		t.Error("test error")
	}
}
