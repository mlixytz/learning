package leetcode392

import "testing"

func TestSolution(t *testing.T) {
	if !isSubsequence("abc", "ahbgdc") {
		t.Error("test error")
	}
	if isSubsequence("axc", "ahbgdc") {
		t.Error("test error")
	}
}
