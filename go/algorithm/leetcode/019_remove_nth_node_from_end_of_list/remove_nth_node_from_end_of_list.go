package leetcode019

/*
   slow-fast pointers
   fast节点先走n步，再和slow一起走，
   当fast到达最后一个结点的时。slow就是倒数n+1个节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, head
	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
