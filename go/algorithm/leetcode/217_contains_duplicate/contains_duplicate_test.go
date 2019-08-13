package leetcode217

import "testing"

func TestSolution(t *testing.T) {
	if !containsDuplicate([]int{1, 2, 3, 1}) {
		t.Error("test error")
	}
	if containsDuplicate([]int{1, 2, 3, 4}) {
		t.Error("test error")
	}
}
