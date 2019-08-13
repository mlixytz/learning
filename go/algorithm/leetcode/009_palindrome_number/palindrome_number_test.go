package leetcode009

import "testing"

func TestSolution(t *testing.T) {
	if isPalindrome(123) {
		t.Error("test erro")
	}
	if !isPalindrome(121) {
		t.Error("test erro")
	}
}
