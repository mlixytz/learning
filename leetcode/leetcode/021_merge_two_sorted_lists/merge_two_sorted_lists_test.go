package leetcode021

import (
	"testing"
)

func TestSolution(t *testing.T) {
	l1 := createList([]int{1, 2, 4})
	l2 := createList([]int{1, 3, 4})
	if !isEqual(mergeTwoLists(l1, l2), createList([]int{1, 1, 2, 3, 4, 4})) {
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

func isEqual(a, b *ListNode) bool {
	for {
		if (a.Next == nil && b.Next != nil) || (a.Next != nil && b.Next == nil) {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		if a.Next == nil {
			return true
		}
		a = a.Next
		b = b.Next
	}
}
