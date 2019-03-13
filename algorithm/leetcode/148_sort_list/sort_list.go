package leetcode148

/*
   链表使用归并排序。时间复杂度为O(nlgn)
   链表分两段的方法是使用 fast-slow 法，用两个指针，
   一个每次走两步，一个走一步，直到fast的走到了末尾，然
   后慢的所在位置就是中间位置，这样就分成了两段
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 将链表分成两段
	slow, fast := head, head
	var mid = new(ListNode)
	for fast != nil && fast.Next != nil {
		mid = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid.Next = nil
	l1 := sortList(head)
	l2 := sortList(slow)
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	var head = new(ListNode)
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}
