package leetcode019

import "testing"

func TestSolution(t *testing.T) {
	if !isEqual(removeNthFromEnd(createList([]int{1, 2, 3, 4, 5}), 2), createList([]int{1, 2, 3, 5})) {
		t.Error("test error")
	}

	if !isEqual(removeNthFromEnd(createList([]int{1}), 1), nil) {
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
		if a == nil && b == nil {
			return true
		}
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
