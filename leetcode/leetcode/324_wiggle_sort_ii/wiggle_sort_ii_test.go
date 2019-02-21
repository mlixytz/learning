package leetcode324

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	if !isEqual(nums, []int{1, 5, 1, 6, 1, 4}) {
		fmt.Println(nums)
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
