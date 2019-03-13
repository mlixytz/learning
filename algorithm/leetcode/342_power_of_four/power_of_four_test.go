package leetcode342

import "testing"

func TestSolution(t *testing.T) {
	if !isPowerOfFour(16) {
		t.Error("test error")
	}

	if isPowerOfFour(5) {
		t.Error("test error")
	}
}
