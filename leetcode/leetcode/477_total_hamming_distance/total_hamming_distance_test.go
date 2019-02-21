package leetcode477

import "testing"

func TestSolution(t *testing.T) {
	if totalHammingDistance([]int{4, 14, 2}) != 6 {
		t.Error("test error")
	}
}
