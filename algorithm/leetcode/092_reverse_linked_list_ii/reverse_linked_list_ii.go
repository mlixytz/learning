package leetcode092

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m >= n {
		return head
	}

	dummyHead := &ListNode{Next: head}

	position := 1

	pre := dummyHead

	// 寻找 m 的前一个节点
	for position < m && head != nil {
		pre = pre.Next
		head = head.Next
		position++
	}

	if head == nil {
		return dummyHead.Next
	}

	// 寻找 n 节点
	for position < n && head.Next != nil {
		head = head.Next
		position++
	}
	next := head.Next

	for pre.Next != head {
		tmp := pre.Next
		pre.Next = pre.Next.Next
		tmp.Next = next
		head.Next = tmp
		next = tmp
	}

	return dummyHead.Next
}
