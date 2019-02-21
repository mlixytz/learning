package leetcode206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	next := head
	for next != nil {
		prev, next.Next, next = next, prev, next.Next
	}
	return prev
}
