package leetcode318

import "testing"

func TestSolution(t *testing.T) {
	if maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}) != 16 {
		t.Error("test error")
	}
	if maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}) != 4 {
		t.Error("test error")
	}
	if maxProduct([]string{"a", "aa", "aaa", "aaaa"}) != 0 {
		t.Error("test error")
	}
}
