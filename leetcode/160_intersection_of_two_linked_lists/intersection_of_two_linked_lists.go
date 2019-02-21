package leetcode160

/*
	如果链表长度相同，循环一次，就可以找到相交节点
	如果链表长度不相同，循环两次，第二次交换链表。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}
