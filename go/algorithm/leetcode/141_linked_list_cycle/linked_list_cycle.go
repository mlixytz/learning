package leetcode141

/*
	使用slow-fast pointers 如果两个指针相遇 就说明链表有环
	证明（slow fast x 都表示元素）：
	    slow == fast   前一步 fast slow
	    fast slow      前一步 fast x slow
	    fast x slow    前一步 fast x x slow
	    fast x x slow  前一步 fast x x x slow
	    ...
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
