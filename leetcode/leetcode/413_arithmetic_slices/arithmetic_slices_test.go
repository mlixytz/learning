package leetcode413

import "testing"

func TestSolution(t *testing.T) {
	if numberOfArithmeticSlices([]int{1, 2, 3, 4}) != 3 {
		t.Error("test error")
	}
}
