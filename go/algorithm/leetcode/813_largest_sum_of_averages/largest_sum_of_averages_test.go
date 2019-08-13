package leetcode813

import "testing"

func TestSolution(t *testing.T) {
	if largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3) != 20 {
		t.Error("test error")
	}
}
