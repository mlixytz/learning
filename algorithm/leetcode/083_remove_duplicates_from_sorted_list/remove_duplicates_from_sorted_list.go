package leetcode083

/*
   当前元素与下一个元素比较（而不是之前的元素）
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	for tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}
