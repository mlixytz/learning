package leetcode147

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Println() string {
	var res string
	iter := ln
	for iter != nil {
		res = fmt.Sprintf("%s%d->", res, iter.Val)
		iter = iter.Next
	}
	return res[:len(res)-2]
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := head
	cur := head.Next
	for cur != nil {
		iter := head
		var prev *ListNode
		for iter != cur {
			if cur.Val < iter.Val {
				// 左边界
				last.Next = cur.Next
				cur.Next = iter
				if prev == nil {
					head = cur
				} else {
					prev.Next = cur
				}
				break
			}
			prev = iter
			iter = iter.Next
		}
		if iter == last.Next {
			cur = last.Next.Next
			last = last.Next
		} else {
			cur = last.Next
		}
	}
	return head
}
