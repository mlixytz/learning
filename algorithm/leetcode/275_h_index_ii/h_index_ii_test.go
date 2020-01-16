package leetcode275

import "testing"

func TestSolution(t *testing.T) {
	if hIndex([]int{0, 1, 3, 5, 6}) != 3 {
		t.Error("test error")
	}
	if hIndex([]int{}) != 0 {
		t.Error("test error")
	}
	if hIndex([]int{100}) != 1 {
		t.Error("test error")
	}
	if hIndex([]int{1, 1}) != 1 {
		t.Error("test error")
	}
	if hIndex([]int{1, 1, 2, 3, 4, 5, 7}) != 3 {
		t.Error("test error")
	}

}
