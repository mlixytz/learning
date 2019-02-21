package leetcode367

import "testing"

func TestSolution(t *testing.T) {
	if isPerfectSquare(8) != false {
		t.Error("test error")
	}
	if isPerfectSquare(16) != true {
		t.Error("test error")
	}
	if isPerfectSquare(1) != true {
		t.Error("test error")
	}
}
