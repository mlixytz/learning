package leetcode764

import "testing"

func TestSolution(t *testing.T) {
	if orderOfLargestPlusSign(5, [][]int{{4, 2}}) != 2 {
		t.Error("test error")
	}
	if orderOfLargestPlusSign(2, [][]int{}) != 1 {
		t.Error("test error")
	}
	if orderOfLargestPlusSign(1, [][]int{{0, 0}}) != 0 {
		t.Error("test error")
	}
	if orderOfLargestPlusSign(2, [][]int{{0, 1}, {1, 0}, {1, 1}}) != 1 {
		t.Error("test error")
	}

}
