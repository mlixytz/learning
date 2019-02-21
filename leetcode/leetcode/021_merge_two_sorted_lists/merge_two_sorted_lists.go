package leetcode021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	last := head

	for l1 != nil || l2 != nil {
		if l1 == nil {
			last.Next = l2
			return head.Next
		}
		if l2 == nil {
			last.Next = l1
			return head.Next
		}
		if l1.Val > l2.Val {
			last.Next = l2
			last = last.Next
			l2 = l2.Next
		} else {
			last.Next = l1
			last = last.Next
			l1 = l1.Next
		}
	}
	return head.Next
}
