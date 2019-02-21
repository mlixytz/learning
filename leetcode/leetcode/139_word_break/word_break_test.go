package leetcode139

import "testing"

func TestSolution(t *testing.T) {
	if !wordBreak("leetcode", []string{"leet", "code"}) {
		t.Error("test error")
	}
	if !wordBreak("applepenapple", []string{"apple", "pen"}) {
		t.Error("test error")
	}
	if wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) {
		t.Error("test error")
	}
}
