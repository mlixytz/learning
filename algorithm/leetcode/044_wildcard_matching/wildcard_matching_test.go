package leetcode044

import "testing"

func TestSolution(t *testing.T) {
	if !isMatch("aa", "*") {
		t.Error("test error")
	}
	if isMatch("cb", "?a") {
		t.Error("test error")
	}
	if !isMatch("adceb", "*a*b") {
		t.Error("test error")
	}
}
