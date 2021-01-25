package leetcode448

import (
	"testing"
)

func TestSolution(t *testing.T) {
	res := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	if res[0] != 5 || res[1] != 6 {
		t.Error("test error")
	}
}
