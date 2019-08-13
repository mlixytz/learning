package leetcode024

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Next: head}
	prev := dummyHead

	cur := head
	for cur != nil && cur.Next != nil {
		tmp := cur.Next.Next
		prev.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = tmp
		prev = cur
		cur = tmp
	}
	return dummyHead.Next
}
