package leetcode898

import "testing"

func TestSolution(t *testing.T) {
	if subarrayBitwiseORs([]int{0}) != 1 {
		t.Error("test error")
	}
	if subarrayBitwiseORs([]int{1, 1, 2}) != 3 {
		t.Error("test error")
	}
	if subarrayBitwiseORs([]int{1, 2, 4}) != 6 {
		t.Error("test error")
	}
}
