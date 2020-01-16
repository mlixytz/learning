package leetcode817

import "testing"

func TestSolution(t *testing.T) {
	if numComponents(createList([]int{0, 1, 2, 3}), []int{0, 1, 3}) != 2 {
		t.Error("test error")
	}
	if numComponents(createList([]int{0, 1, 2, 3, 4}), []int{0, 3, 1, 4}) != 2 {
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
