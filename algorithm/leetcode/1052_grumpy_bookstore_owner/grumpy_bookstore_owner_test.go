package leetcode1052

import "testing"

func TestSolution(t *testing.T) {
	if maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3) != 16 {
		t.Error("test error")
	}
}
