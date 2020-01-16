package leetcode389

import "testing"

func TestSolution(t *testing.T) {
	if findTheDifference("abcd", "abcde") != byte('e') {
		t.Error("test error")
	}
}
