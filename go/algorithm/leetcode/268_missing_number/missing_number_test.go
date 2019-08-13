package leetcode268

import "testing"

func TestSolution(t *testing.T) {
	if missingNumber([]int{3, 0, 1}) != 2 {
		t.Error("test error")
	}
	if missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) != 8 {
		t.Error("test error")
	}
}
