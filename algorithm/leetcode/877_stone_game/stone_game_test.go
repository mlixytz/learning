package leetcode877

import "testing"

func TestSolution(t *testing.T) {
	if stoneGame([]int{5, 3, 4, 5}) != true {
		t.Error("test error")
	}
}
