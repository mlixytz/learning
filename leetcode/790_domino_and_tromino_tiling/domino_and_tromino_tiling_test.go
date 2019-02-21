package leetcode790

import "testing"

func TestSolution(t *testing.T) {
	if numTilings(3) != 5 {
		t.Error("test error")
	}
	if numTilings(4) != 11 {
		t.Error("test error")
	}
	if numTilings(30) != 9312342245 {
		t.Error("test error")
	}
}
