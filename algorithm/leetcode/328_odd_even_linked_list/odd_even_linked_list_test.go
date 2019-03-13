package leetcode328

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if !isEqual(oddEvenList(createList([]int{1, 2, 3, 4, 5})), createList([]int{1, 3, 5, 2, 4})) {
		t.Error("test error")
	}
	if !isEqual(oddEvenList(createList([]int{2, 1, 3, 5, 6, 4, 7})), createList([]int{2, 3, 6, 7, 1, 5, 4})) {
		t.Error("test error")
	}
	if !isEqual(oddEvenList(createList([]int{1, 2, 3, 4})), createList([]int{1, 3, 2, 4})) {
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
