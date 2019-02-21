package leetcode081

import "testing"

func TestSolution(t *testing.T) {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	if search(nums, 0) != true {
		t.Error("test error")
	}
	if search(nums, 3) != false {
		t.Error("test error")
	}
	nums = []int{1, 1, 3, 1}
	if search(nums, 3) != true {
		t.Error("test error")
	}
}

func TestBestSolution(t *testing.T) {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	if searchBestSolution(nums, 0) != true {
		t.Error("test error")
	}
	if searchBestSolution(nums, 3) != false {
		t.Error("test error")
	}
	nums = []int{1, 1, 3, 1}
	if searchBestSolution(nums, 3) != true {
		t.Error("test error")
	}
}
