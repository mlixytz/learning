package leetcode179

import "testing"

func TestSolution(t *testing.T) {
	if largestNumber([]int{10, 2}) != "210" {
		t.Error("test error")
	}
	if largestNumber([]int{3, 30, 34, 5, 9}) != "9534330" {
		t.Error("test error")
	}
	if largestNumber([]int{3353, 335}) != "3353353" {
		t.Error("test error")
	}
	if largestNumber([]int{0, 0}) != "0" {
		t.Error("test error")
	}
}
