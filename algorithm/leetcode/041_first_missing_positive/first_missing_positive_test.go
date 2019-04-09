package leetcode041

import "testing"

func TestSolution(t *testing.T) {
	if firstMissingPositive([]int{3, 4, -1, 1}) != 2 {
		t.Error("test error")
	}
}
