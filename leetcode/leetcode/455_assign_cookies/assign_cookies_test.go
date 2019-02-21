package leetcode455

import "testing"

func TestSolution(t *testing.T) {
	if findContentChildren([]int{1, 2, 3}, []int{1, 1}) != 1 {
		t.Error("test error")
	}
	if findContentChildren([]int{1, 2}, []int{1, 2, 3}) != 2 {
		t.Error("test error")
	}
}
