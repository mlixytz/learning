package leetcode025

import (
	"testing"

	"github.com/mlixytz/learning/algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	l1 := utils.CreateSingleLinkedList([]int{1, 2, 3, 4, 5})
	l2 := utils.CreateSingleLinkedList([]int{2, 1, 4, 3, 5})
	if !utils.TwoSingleLinkedListIsEqual(reverseKGroup(l1, 2), l2) {
		t.Error("test error")
	}
}
