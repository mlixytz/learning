package leetcode080

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if removeDuplicates([]int{1, 1, 1, 2, 2, 3}) != 5 {
		t.Error("test error")
	}
	if removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}) != 7 {
		t.Error("test error")
	}
}
