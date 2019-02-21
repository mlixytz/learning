package leetcode028

import "testing"

func TestSolution(t *testing.T) {
	if strStr("hello", "ll") != 2 {
		t.Error("test error")
	}
	if strStr("aaaaa", "bba") != -1 {
		t.Error("test error")
	}
	if strStr("hlello", "ll") != 3 {
		t.Error("test error")
	}
	if strStr("mississippi", "issip") != 4 {
		t.Error("test error")
	}
}
