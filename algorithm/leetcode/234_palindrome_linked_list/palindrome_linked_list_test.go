package leetcode234

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if isPalindrome(createList([]int{1, 2})) {
		t.Error("test error")
	}
	if !isPalindrome(createList([]int{1, 2, 2, 1})) {
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
