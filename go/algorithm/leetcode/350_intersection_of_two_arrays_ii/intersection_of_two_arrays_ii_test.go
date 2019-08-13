package leetcode350

import "testing"

func TestSolution(t *testing.T) {
	if !isEqual(intersect([]int{1, 2, 2, 1}, []int{2, 2}), []int{2, 2}) {
		t.Error("test error")
	}
	if !isEqual(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{9, 4}) && !isEqual(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{4, 9}) {
		t.Error("test error")
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != a[i] {
			return false
		}
	}
	return true
}
