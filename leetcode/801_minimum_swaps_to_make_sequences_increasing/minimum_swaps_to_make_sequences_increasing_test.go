package leetcode801

import "testing"

func TestSolution(t *testing.T) {
	if minSwap([]int{1, 3, 5, 4}, []int{1, 2, 3, 7}) != 1 {
		t.Error("test error")
	}
}
