package leetcode121

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if maxProfit([]int{7, 1, 5, 3, 6, 4}) != 5 {
		t.Error("test error")
	}
	if maxProfit([]int{7, 6, 4, 3, 1}) != 0 {
		t.Error("test error")
	}
	if maxProfit([]int{}) != 0 {
		t.Error("test error")
	}
	if maxProfit([]int{1}) != 0 {
		t.Error("test error")
	}
}
