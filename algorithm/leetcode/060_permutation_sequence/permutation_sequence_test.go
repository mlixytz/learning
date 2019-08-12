package leetcode060

import "testing"

func TestSolution(t *testing.T) {
	if getPermutation(4, 5) != "1423" {
		t.Error("test error")
	}
	if getPermutation(2, 1) != "12" {
		t.Error("test error")
	}
	if getPermutation(2, 2) != "21" {
		t.Error("test error")
	}
	if getPermutation(1, 1) != "1" {
		t.Error("test error")
	}
	if getPermutation(3, 2) != "132" {
		t.Error("test error")
	}
}
