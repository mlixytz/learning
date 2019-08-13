package leetcode714

import "testing"

func TestSolution(t *testing.T) {
	if maxProfit([]int{1, 3, 2, 8, 4, 9}, 2) != 8 {
		t.Error("test error")
	}
}
