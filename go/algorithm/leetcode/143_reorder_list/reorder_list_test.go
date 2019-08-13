package leetcode143

import "testing"

func TestSolution(t *testing.T) {
	ln := createList([]int{1, 2, 3, 4})
	reorderList(ln)
	if !isEqual(createList([]int{1, 4, 2, 3}), ln) {
		t.Error("test error")
	}
	ln2 := createList([]int{1, 2, 3, 4, 5})
	reorderList(ln2)
	if !isEqual(createList([]int{1, 5, 2, 4, 3}), ln2) {
		t.Error("test error")
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
