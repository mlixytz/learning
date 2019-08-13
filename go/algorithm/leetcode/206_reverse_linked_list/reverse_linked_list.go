package leetcode206

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
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

// 递归
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	return reverse(nil, head)
// }

// func reverse(prev, cur *ListNode) *ListNode {
// 	next := cur.Next
// 	cur.Next = prev
// 	if next != nil {
// 		return reverse(cur, next)
// 	}
// 	return cur
// }
