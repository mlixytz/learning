package leetcode935

import "testing"

func TestSolution(t *testing.T) {
	if knightDialer(1) != 10 {
		t.Error("test error")
	}
	if knightDialer(2) != 20 {
		t.Error("test error")
	}
	if knightDialer(3) != 46 {
		t.Error("test error")
	}
}
