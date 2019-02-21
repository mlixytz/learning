package leetcode234

/*
   用slow-fast pointers 找到中间元素，然后反转一半列表与未反转的列表对比
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil { // 奇数
		slow = slow.Next
	}
	revHead := reverse(*slow)
	for revHead != nil {
		if revHead.Val != head.Val {
			return false
		}
		revHead, head = revHead.Next, head.Next
	}
	return true
}

func reverse(head ListNode) *ListNode {
	if &head == nil && &head.Next == nil {
		return &head
	}

	var prev *ListNode
	cur := &head
	for cur != nil {
		t := cur.Next
		cur.Next = prev
		prev = cur
		cur = t
	}
	return prev
}
