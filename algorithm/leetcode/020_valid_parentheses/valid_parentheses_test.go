package leetcode020

import "testing"

func TestSolution(t *testing.T) {
	if !isValid("()") {
		t.Error("test error")
	}
	if !isValid("()[]{}") {
		t.Error("test error")
	}
	if isValid("(]") {
		t.Error("test error")
	}
	if !isValid("{[]}") {
		t.Error("test error")
	}
}
