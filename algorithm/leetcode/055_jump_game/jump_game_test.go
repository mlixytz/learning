package leetcode055

import "testing"

func TestSolution(t *testing.T) {
	if !canJump([]int{2, 3, 1, 1, 4}) {
		t.Error("test error")
	}
	if canJump([]int{3, 2, 1, 0, 4}) {
		t.Error("test error")
	}
	if canJump([]int{0, 1}) {
		t.Error("test error")
	}
	if !canJump([]int{1, 0}) {
		t.Error("test error")
	}
	if !canJump([]int{1}) {
		t.Error("test error")
	}
}
