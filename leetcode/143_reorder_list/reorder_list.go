package leetcode143

/*
   fast-slow pointers 找到中间结点，再对链表后半部分反转，
   然后合并前半部分和反转的后半部分
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	slow, fast := head, head
	var t *ListNode
	for fast != nil && fast.Next != nil {
		t = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		tmp := slow.Next
		slow.Next = nil
		slow = tmp
	} else {
		t.Next = nil
	}

	// 翻转链表后半部分
	var prev *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}

	// 链表前后部分合并
	f := head
	for prev != nil {
		tmp := prev.Next
		tmpf := f.Next
		f.Next = prev
		f = tmpf
		prev.Next = f
		prev = tmp
	}
}
