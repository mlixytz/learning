package leetcode876

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if middleNode(createList([]int{1, 2, 3, 4, 5})).Val != 3 {
		t.Error("test error")
	}
	if middleNode(createList([]int{1, 2, 3, 4, 5, 6})).Val != 4 {
		t.Error("test error")
	}
}

func createList(nums []int) *ListNode {
	head := &ListNode{}
	tmp := head
	for i, num := range nums {
		tmp.Val = num
		if i < len(nums)-1 {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}
	}
	return head
}
