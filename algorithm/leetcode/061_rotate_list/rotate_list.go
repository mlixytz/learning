package leetcode061

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 1

	cur := head
	for cur.Next != nil {
		length++
		cur = cur.Next
	}

	position := k % length
	if position == 0 {
		return head
	}

	cur.Next = head

	position = length - position

	split := head

	index := 1

	for split != nil {
		head = split.Next
		if index == position {
			split.Next = nil
			break
		}
		split = split.Next
		index++
	}
	return head
}
