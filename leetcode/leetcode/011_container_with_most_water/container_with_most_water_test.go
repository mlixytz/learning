package leetcode011

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) != 49 {
		t.Error("test error")
	}
}
