package leetcode125

import "testing"

func TestSolution(t *testing.T) {
	if !isPalindrome("A man, a plan, a canal: Panama") {
		t.Error("test error")
	}
	if isPalindrome("race a car") {
		t.Error("test error")
	}
	if isPalindrome("0P") {
		t.Error("test error")
	}

}
