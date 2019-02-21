package leetcode086

/*
   新创建两个链表，一个小于x，一个大于x， 最后合并两个链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	less := &ListNode{}
	geater := &ListNode{}
	lessHead := less
	geaterHead := geater
	var tmp *ListNode
	for head != nil {
		tmp = head.Next
		head.Next = nil
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			geater.Next = head
			geater = geater.Next
		}
		head = tmp
	}
	less.Next = geaterHead.Next
	return lessHead.Next
}
