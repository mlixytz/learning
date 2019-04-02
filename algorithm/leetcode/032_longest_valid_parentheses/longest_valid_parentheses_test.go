package leetcode032

import "testing"

func TestSolution(t *testing.T) {
	if longestValidParentheses("(()") != 2 {
		t.Error("test erorr")
	}
	if longestValidParentheses("())") != 2 {
		t.Error("test erorr")
	}
}
