package leetcode007

import "testing"

func TestSolution(t *testing.T) {
	if reverse(123) != 321 {
		t.Error("test error")
	}
	if reverse(-123) != -321 {
		t.Error("test error")
	}
}
