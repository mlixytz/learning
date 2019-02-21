package leetcode169

import "testing"

func TestSolution(t *testing.T) {
	if majorityElement([]int{3, 2, 3}) != 3 {
		t.Error("test error")
	}
	if majorityElement([]int{2, 2, 1, 1, 1, 2, 2}) != 2 {
		t.Error("test error")
	}
}
