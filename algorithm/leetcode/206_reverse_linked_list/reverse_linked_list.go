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

// 递归实现
// func reverseList(head *ListNode) *ListNode {
// 	return reverse(nil, head)
// }

// func reverse(prev *ListNode, cur *ListNode) *ListNode{
// 	if cur == nil {
// 			return prev
// 	}
// 	next := cur.Next
// 	cur.Next = prev
// 	return reverse(cur, next)
// }
