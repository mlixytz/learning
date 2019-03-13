package leetcode300

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}) != 4 {
		t.Error("test error")
	}
	if lengthOfLIS([]int{10, 9, 2, 5, 3, 4}) != 3 {
		t.Error("test error")
	}
	if lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}) != 6 {
		t.Error("test error")
	}

	if lengthOfLISBS([]int{10, 9, 2, 5, 3, 7, 101, 18}) != 4 {
		t.Error("test error")
	}
	if lengthOfLISBS([]int{10, 9, 2, 5, 3, 4}) != 3 {
		t.Error("test error")
	}
	if lengthOfLISBS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}) != 6 {
		t.Error("test error")
	}

}
