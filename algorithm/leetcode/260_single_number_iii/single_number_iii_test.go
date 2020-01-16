package leetcode260

import "testing"

func TestSolution(t *testing.T) {
	if !isEqual(singleNumber([]int{1, 2, 1, 3, 2, 5}), []int{3, 5}) {
		t.Error("test error")
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
