package leetcode725

import "testing"

func TestSolution(t *testing.T) {
	ln := splitListToParts(createList([]int{1, 2, 3, 4}), 5)
	correct := []*ListNode{createList([]int{1}), createList([]int{2}), createList([]int{3}), createList([]int{4}), nil}
	for i, val := range correct {
		if !isEqual(val, ln[i]) {
			t.Error("test error")
		}
	}

	ln2 := splitListToParts(createList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 3)
	correct2 := []*ListNode{createList([]int{1, 2, 3, 4}), createList([]int{5, 6, 7}), createList([]int{8, 9, 10})}
	for i, val := range correct2 {
		if !isEqual(val, ln2[i]) {
			t.Error("test error")
		}
	}
}

func createList(nums []int) *ListNode {
	head := &ListNode{}
	tmp := head
	for i, num := range nums {
		tmp.Val = num
		if i < len(nums)-1 {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}
	}
	return head
}

func isEqual(a, b *ListNode) bool {
	for {
		if a == nil && b == nil {
			return true
		}
		if (a.Next == nil && b.Next != nil) || (a.Next != nil && b.Next == nil) {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		if a.Next == nil {
			return true
		}
		a = a.Next
		b = b.Next
	}
}
