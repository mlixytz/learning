package leetcode042

import "testing"

func TestSolution(t *testing.T) {
	if trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) != 6 {
		t.Error("test error")
	}
}
