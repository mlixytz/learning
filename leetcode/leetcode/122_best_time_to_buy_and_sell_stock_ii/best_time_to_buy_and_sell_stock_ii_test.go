package leetcode122

import "testing"

func TestSolution(t *testing.T) {
	if maxProfit([]int{7, 1, 5, 3, 6, 4}) != 7 {
		t.Error("test error")
	}
	if maxProfit([]int{1, 2, 3, 4, 5}) != 4 {
		t.Error("test error")
	}
	if maxProfit([]int{7, 6, 4, 3, 1}) != 0 {
		t.Error("test error")
	}
}
