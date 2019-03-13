package leetcode023

import "github.com/mlixytz/go-algorithm/leetcode/utils"

type ListNode = utils.SingleLinkedListNode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else {
		return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	head := &ListNode{}
	last := head

	for l1 != nil || l2 != nil {
		if l1 == nil {
			last.Next = l2
			return head.Next
		}
		if l2 == nil {
			last.Next = l1
			return head.Next
		}
		if l1.Val > l2.Val {
			last.Next = l2
			last = last.Next
			l2 = l2.Next
		} else {
			last.Next = l1
			last = last.Next
			l1 = l1.Next
		}
	}
	return head.Next
}
