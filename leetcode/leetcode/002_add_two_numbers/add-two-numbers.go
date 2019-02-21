package leetcode002

// ListNode is a node of Singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := &ListNode{}
	currNode := headNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		currNode.Next = &ListNode{Val: sum % 10}
		currNode = currNode.Next
	}
	if carry > 0 {
		currNode.Next = &ListNode{Val: carry}
	}
	return headNode.Next
}
