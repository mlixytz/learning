package leetcode237

import "testing"

func TestSolution(t *testing.T) {
	head := createList([]int{4, 5, 1, 9})
	deleteNode(head.Next)
	if !isEqual(head, createList([]int{4, 1, 9})) {
		t.Error("test error")
	}

	head = createList([]int{4, 5, 1, 9})
	deleteNode(head.Next.Next)
	if !isEqual(head, createList([]int{4, 5, 9})) {
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
