package leetcode136

import "testing"

func TestSolution(t *testing.T) {
	if singleNumber([]int{2, 2, 1}) != 1 {
		t.Error("test error")
	}
	if singleNumber([]int{4, 1, 2, 1, 2}) != 4 {
		t.Error("test error")
	}
}
