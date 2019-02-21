package leetcode026

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) != 5 {
		t.Error("test error")
	}

}
