package leetcode524

import "testing"

func TestSolution(t *testing.T) {
	if findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}) != "apple" {
		t.Error("test error")
	}
	if findLongestWord("abpcplea", []string{"a", "b", "c"}) != "a" {
		t.Error("test error")
	}
	if findLongestWord("bab", []string{"ba", "ab", "c"}) != "ab" {
		t.Error("test error")
	}
}
