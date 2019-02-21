package leetcode740

import "testing"

func TestSolution(t *testing.T) {
	if deleteAndEarn([]int{3, 4, 2}) != 6 {
		t.Error("test error")
	}
	if deleteAndEarn([]int{2, 2, 3, 3, 3, 4}) != 9 {
		t.Error("test error")
	}
}
