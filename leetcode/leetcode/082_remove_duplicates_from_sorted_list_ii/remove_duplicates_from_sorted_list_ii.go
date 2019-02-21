package leetcode082

/*
   定义假head，pre, cur
   每次判断 cur.Val 和 cur.Next.Val 的值，来确实是否重复
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyhead := &ListNode{Next: head}
	pre := dummyhead
	cur := dummyhead.Next
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return dummyhead.Next
}
