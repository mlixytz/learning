package leetcode148

import "testing"

func TestSolution(t *testing.T) {
	l1 := makeList([]int{4, 2, 1, 3})
	l2 := makeList([]int{1, 2, 3, 4})
	l3 := makeList([]int{-1, 5, 3, 4, 0})
	l4 := makeList([]int{-1, 0, 3, 4, 5})
	if !isEqual(sortList(l1), l2) {
		t.Error("test error")
	}
	if !isEqual(sortList(l3), l4) {
		t.Error("test error")
	}
}

func makeList(nums []int) *ListNode {
	var head = new(ListNode)
	cur := head
	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return head.Next
}

func isEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		return false
	}
	return true
}
