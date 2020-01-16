package leetcode058

import "testing"

func TestSolution(t *testing.T) {
	if lengthOfLastWord("Hello World") != 5 {
		t.Error("test error")
	}
	if lengthOfLastWord("") != 0 {
		t.Error("test error")
	}
}
