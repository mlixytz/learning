package leetcode516

import "testing"

func TestSolution(t *testing.T) {
	if longestPalindromeSubseq("bbbab") != 4 {
		t.Error("test error")
	}
	if longestPalindromeSubseq("cbbd") != 2 {
		t.Error("test error")
	}
}
