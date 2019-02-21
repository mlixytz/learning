package leetcode338

import "testing"

func TestSolution(t *testing.T) {
	if !isEqual(countBits(2), []int{0, 1, 1}) {
		t.Error("test error")
	}
	if !isEqual(countBits(5), []int{0, 1, 1, 2, 1, 2}) {
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
