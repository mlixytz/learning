package leetcode405

import "testing"

func TestSolution(t *testing.T) {
	if toHex(26) != "1a" {
		t.Error("test error")
	}
	if toHex(-1) != "ffffffff" {
		t.Error("test error")
	}
}
