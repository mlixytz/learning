package leetcode033

import "testing"

func TestSolution(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	if search(nums, 0) != 4 {
		t.Error("test error")
	}
	if search(nums, 3) != -1 {
		t.Error("test error")
	}
	if search([]int{}, 5) != -1 {
		t.Error("test error")
	}
	if search([]int{5, 1, 3}, 5) != 0 {
		t.Error("test error")
	}

}
