package leetcode445

import "container/list"

/*
   将链表节点存储到container/list中, 按位运算
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := list.New()
	list2 := list.New()

	for l1 != nil {
		list1.PushFront(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		list2.PushFront(l2.Val)
		l2 = l2.Next
	}

	var prev *ListNode
	var addone, v1, v2, val int
	for e1, e2 := list1.Front(), list2.Front(); e1 != nil || e2 != nil || addone != 0; {
		if e1 == nil {
			v1 = 0
		} else {
			v1 = e1.Value.(int)
			e1 = e1.Next()
		}

		if e2 == nil {
			v2 = 0
		} else {
			v2 = e2.Value.(int)
			e2 = e2.Next()
		}

		if (v1 + v2 + addone) >= 10 {
			val = (v1 + v2 + addone) % 10
			addone = 1
		} else {
			val = v1 + v2 + addone
			addone = 0
		}
		prev = &ListNode{val, prev}
	}

	return prev
}
