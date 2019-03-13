package utils

// SingleLinkedListNode 是一个单链表节点
type SingleLinkedListNode struct {
	Val  int
	Next *SingleLinkedListNode
}

// CreateSingleLinkedList 创建单链表
func CreateSingleLinkedList(nums []int) *SingleLinkedListNode {
	head := &SingleLinkedListNode{}
	tmp := head
	for i, num := range nums {
		tmp.Val = num
		if i < len(nums)-1 {
			tmp.Next = &SingleLinkedListNode{}
			tmp = tmp.Next
		}
	}
	return head
}

// TwoSingleLinkedListIsEqual 两个但链表相等
func TwoSingleLinkedListIsEqual(a, b *SingleLinkedListNode) bool {
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
