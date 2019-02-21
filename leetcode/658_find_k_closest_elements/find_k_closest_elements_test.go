package leetcode658

import (
	"testing"
)

func TestSolution(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !compareTwoSlice(findClosestElements(arr, 4, 3), []int{1, 2, 3, 4}) {
		t.Error("test error")
	}
	if !compareTwoSlice(findClosestElements(arr, 4, -1), []int{1, 2, 3, 4}) {
		t.Error("test error")
	}
	if !compareTwoSlice(findClosestElements(arr, 4, 6), []int{2, 3, 4, 5}) {
		t.Error("test error")
	}
}

func compareTwoSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}
