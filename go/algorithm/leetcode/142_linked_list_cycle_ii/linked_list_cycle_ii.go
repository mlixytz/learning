package leetcode142

/*
    slow-fast pointers （slow 一次一步， fast 一次2步）
    如果有环 两个指针第一次相交的的位置 到链表最后一个节点的距离等于 节点head到循环开始的距离
    证明：
	设链表head到循环开始的距离为 X1
	开始循环到第一次相交的距离为 X2
	第一次相交到链表tail的距离为 X3

	于是 x1 + x2 + x3 + x2 = 2(x1+x2)
		    ||
		 x1 == x3
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return slow
		}
	}
	return nil
}
