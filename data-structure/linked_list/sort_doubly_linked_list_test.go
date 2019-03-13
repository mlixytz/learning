package linkedlist

import (
	"container/list"
	"testing"
)

func TestSortDoublyLinkedList(t *testing.T) {
	l1 := list.New()
	l1.PushBack(4)
	l1.PushBack(3)
	l1.PushBack(7)
	l1.PushBack(10)
	l1.PushBack(1)
	l1.PushBack(9)
	l1.PushBack(8)

	l1 = QuickSort(l1)
	nums := []int{1, 3, 4, 7, 8, 9, 10}
	for i, j := l1.Front(), 0; i != nil; i = i.Next() {
		if i.Value.(int) != nums[j] {
			t.Error("quick sort error")
		}
		j++
	}
}
