package leetcode050

import "testing"

func TestSolution(t *testing.T) {
	if myPow(3, 3) != float64(27) {
		t.Error("test error")
	}
	if myPow(3, -1) != float64(1)/float64(3) {
		t.Error("test error")
	}

}
