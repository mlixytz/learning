package leetcode393

import "testing"

func TestSolution(t *testing.T) {
	if !validUtf8([]int{197, 130, 1}) {
		t.Error("test error")
	}
	if validUtf8([]int{235, 140, 4}) {
		t.Error("test error")
	}
	if validUtf8([]int{145}) {
		t.Error("test error")
	}
}
