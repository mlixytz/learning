package leetcode083

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if !isEqual(deleteDuplicates(createList([]int{1, 1, 2})), createList([]int{1, 2})) {
		t.Error("test error")
	}
	if !isEqual(deleteDuplicates(createList([]int{1, 1, 2, 3, 3})), createList([]int{1, 2, 3})) {
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
