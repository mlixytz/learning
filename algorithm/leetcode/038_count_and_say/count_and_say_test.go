package leetcode038

import "testing"

func TestSolution(t *testing.T) {
	if countAndSay(5) != "111221" {
		t.Error("test error")
	}
}
