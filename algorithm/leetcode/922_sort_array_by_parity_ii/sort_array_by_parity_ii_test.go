package leetcode922

import "testing"

func TestSolution(t *testing.T) {
	if !isEqual(sortArrayByParityII([]int{4, 2, 5, 7}), []int{4, 5, 2, 7}) && !isEqual(sortArrayByParityII([]int{4, 2, 5, 7}), []int{4, 7, 2, 5}) && !isEqual(sortArrayByParityII([]int{4, 2, 5, 7}), []int{2, 5, 4, 7}) && !isEqual(sortArrayByParityII([]int{4, 2, 5, 7}), []int{2, 7, 4, 5}) {
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
