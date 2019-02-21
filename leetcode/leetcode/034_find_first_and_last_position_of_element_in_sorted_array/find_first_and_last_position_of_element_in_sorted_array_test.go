package leetcode034

import "testing"

func TestSolution(t *testing.T) {
	if !compareTwoSlice(searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4}) {
		t.Error("test Error")
	}
	if !compareTwoSlice(searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1}) {
		t.Error("test Error")
	}
	if !compareTwoSlice(searchRange([]int{2, 2}, 2), []int{0, 1}) {
		t.Error("test Error")
	}
	if !compareTwoSlice(searchRangeBestSolution([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4}) {
		t.Error("test Error")
	}
	if !compareTwoSlice(searchRangeBestSolution([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1}) {
		t.Error("test Error")
	}
	if !compareTwoSlice(searchRangeBestSolution([]int{2, 2}, 2), []int{0, 1}) {
		t.Error("test Error")
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
