package leetcode069

import "testing"

func TestSolution(t *testing.T) {
	if mySqrt(15) != 3 {
		t.Error("test error")
	}
	if mySqrt(10) != 3 {
		t.Error("test error")
	}
	if mySqrt(0) != 0 {
		t.Error("test error")
	}
	if mySqrt(1) != 1 {
		t.Error("test error")
	}

}
