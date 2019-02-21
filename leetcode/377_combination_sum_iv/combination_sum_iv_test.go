package leetcode377

import "testing"

func TestSolution(t *testing.T) {
	if combinationSum4([]int{1, 2, 3}, 4) != 7 {
		t.Error("test error")
	}
}
