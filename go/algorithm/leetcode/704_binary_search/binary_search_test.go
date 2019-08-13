package leetcode704

import "testing"

func TestSolution(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	if search(nums, 9) != 4 {
		t.Error("test Error")
	}
	if search(nums, 2) != -1 {
		t.Error("test Error")
	}
}
