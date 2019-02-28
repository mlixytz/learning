package leetcode008

import "testing"

func TestSolution(t *testing.T) {
	if myAtoi("42") != 42 {
		t.Error("test error")
	}
	if myAtoi("   -42") != -42 {
		t.Error("test error")
	}
	if myAtoi("332 ab") != 332 {
		t.Error("test error")
	}
	if myAtoi("word 1b") != 0 {
		t.Error("test error")
	}
	if myAtoi("-91283472332") != -2147483648 {
		t.Error("test error")
	}
}
