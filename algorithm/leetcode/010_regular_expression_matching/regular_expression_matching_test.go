package leetcode010

import "testing"

func TestSolution(t *testing.T) {
	if !isMatch("aab", "c*a*b") {
		t.Error("test error")
	}
	if !isMatch("", "*") {
		t.Error("test error")
	}
	if isMatch("aaa", "aaaa") {
		t.Error("test error")
	}
}
