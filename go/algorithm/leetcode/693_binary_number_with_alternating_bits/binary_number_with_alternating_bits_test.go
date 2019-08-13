package leetcode693

import "testing"

func TestSolution(t *testing.T) {
	if !hasAlternatingBits(5) {
		t.Error("test error")
	}
	if hasAlternatingBits(7) {
		t.Error("test error")
	}
	if hasAlternatingBits(11) {
		t.Error("test error")
	}
	if !hasAlternatingBits(10) {
		t.Error("test error")
	}
	if hasAlternatingBits(4) {
		t.Error("test error")
	}
}
