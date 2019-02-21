package leetcode287

import "testing"

func TestSolution(t *testing.T) {
	if findDuplicate([]int{1, 3, 4, 2, 2}) != 2 {
		t.Error("test error")
	}
}
