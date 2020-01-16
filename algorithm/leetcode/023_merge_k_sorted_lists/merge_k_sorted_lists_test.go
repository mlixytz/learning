package leetcode023

import (
	"testing"

	"github.com/mlixytz/go-algorithm/leetcode/utils"
)

func TestSolution(t *testing.T) {
	lists := []*ListNode{utils.CreateSingleLinkedList([]int{1, 4, 5}), utils.CreateSingleLinkedList([]int{1, 3, 4}), utils.CreateSingleLinkedList([]int{2, 6})}
	if !utils.TwoSingleLinkedListIsEqual(mergeKLists(lists), utils.CreateSingleLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6})) {
		t.Error("test error")
	}
}
