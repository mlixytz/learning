package leetcode027

import "testing"

func TestSolution(t *testing.T) {
	if removeElement([]int{3, 2, 2, 3}, 3) != 2 {
		t.Error("test error")
	}
}
