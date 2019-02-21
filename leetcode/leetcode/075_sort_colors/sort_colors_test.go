package leetcode075

import "testing"

func TestSolution(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	if !isEqual(nums, []int{0, 0, 1, 1, 2, 2}) {
		t.Error("test error")
	}
	nums = []int{2, 0, 1}
	sortColors(nums)
	if !isEqual(nums, []int{0, 1, 2}) {
		t.Error("test error")
	}
	nums = []int{1, 2, 0}
	sortColors(nums)
	if !isEqual(nums, []int{0, 1, 2}) {
		t.Error("test error")
	}
	if !isEqual([]int{2}, []int{2}) {
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
