package leetcode707

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	i := 0
	head := this.Next
	for head != nil {
		if i == index {
			return head.Val
		}
		head = head.Next
		i++
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	head := this.Next
	this.Next = &MyLinkedList{val, head}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	head := this
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &MyLinkedList{Val: val}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	i := -1
	head := this
	for head != nil {
		if i == index-1 {
			head.Next = &MyLinkedList{val, head.Next}
			return
		}
		head = head.Next
		i++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	i := -1
	head := this
	for head.Next != nil {
		if i == index-1 {
			head.Next = head.Next.Next
			return
		}
		head = head.Next
		i++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
