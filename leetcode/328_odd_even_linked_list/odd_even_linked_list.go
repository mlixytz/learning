package leetcode328

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	eventStart := head.Next
	odd := head
	event := eventStart
	for event != nil && event.Next != nil {
		odd.Next, event.Next = event.Next, event.Next.Next
		odd, event = odd.Next, event.Next
	}
	odd.Next = eventStart
	return head
}
