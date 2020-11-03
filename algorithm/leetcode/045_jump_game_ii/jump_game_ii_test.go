package leetcode045

import "testing"

func TestSolution(t *testing.T) {
	if jump([]int{2, 1}) != 1 {
		t.Error("test error")
	}
	if jump([]int{3, 2, 1}) != 1 {
		t.Error("test error")
	}
	if jump([]int{2, 3, 1, 1, 1}) != 2 {
		t.Error("test error")
	}
}
