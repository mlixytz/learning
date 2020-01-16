package leetcode141

import "testing"

func TestSolution(t *testing.T) {
	if !hasCycle(createCycleList([]int{3, 2, 0. - 4}, 1)) {
		t.Error("test error")
	}
	if hasCycle(createList([]int{1})) {
		t.Error("test error")
	}
}

func createCycleList(nums []int, pos int) *ListNode {
	head := &ListNode{}
	tmp := head
	var cycleStart *ListNode
	for i, num := range nums {
		tmp.Val = num
		if i < len(nums)-1 {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		} else {
			tmp.Next = cycleStart
			break
		}
		if i == pos {
			cycleStart = tmp
		}
	}
	return head
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
