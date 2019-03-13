package leetcode416

import "testing"

func TestSolution(t *testing.T) {
	if canPartition([]int{1, 5, 11, 5}) != true {
		t.Error("test error")
	}
}
