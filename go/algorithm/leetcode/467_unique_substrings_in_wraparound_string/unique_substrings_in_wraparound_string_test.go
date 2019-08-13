package leetcode467

import "testing"

func TestSolution(t *testing.T) {
	if findSubstringInWraproundString("a") != 1 {
		t.Error("test error")
	}
	if findSubstringInWraproundString("cac") != 2 {
		t.Error("test error")
	}
	if findSubstringInWraproundString("zab") != 6 {
		t.Error("test error")
	}
	if findSubstringInWraproundString("za") != 3 {
		t.Error("test error")
	}
}
