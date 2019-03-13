package leetcode844

import "testing"

func TestSolution(t *testing.T) {
	if !backspaceCompare("ab#c", "ad#c") {
		t.Error("test error")
	}
	if backspaceCompare("a#c", "b") {
		t.Error("test error")
	}
	if !backspaceCompare("ab##", "c#d#") {
		t.Error("test error")
	}
	if !backspaceCompare("a##c", "#a#c") {
		t.Error("test error")
	}
}
