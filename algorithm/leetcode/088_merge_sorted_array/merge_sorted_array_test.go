package leetcode088

import (
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	if !utils.IsArrayEqual(nums1, []int{1, 2, 2, 3, 5, 6, 0}) {
		t.Error("test error")
	}
}
