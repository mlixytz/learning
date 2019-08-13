package leetcode659

import "testing"

func TestSolution(t *testing.T) {
	if !isPossible([]int{1, 2, 3, 3, 4, 5}) {
		t.Error("test error")
	}
	if !isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}) {
		t.Error("test error")
	}
	if isPossible([]int{1, 2, 3, 4, 4, 5}) {
		t.Error("test error")
	}
}
