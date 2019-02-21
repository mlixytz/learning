package leetcode274

import "testing"

func TestSolution(t *testing.T) {
	if hIndex([]int{3, 0, 6, 1, 5}) != 3 {
		t.Error("test error")
	}
}
