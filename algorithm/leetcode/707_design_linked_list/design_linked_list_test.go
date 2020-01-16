package leetcode707

import "testing"

func TestSolution(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	if linkedList.Get(1) != 2 {
		t.Error("test error")
	} //返回2
	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	if linkedList.Get(1) != 3 {
		t.Error("test error")
	}
}
