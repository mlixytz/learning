package leetcode035

import "testing"

func TestSolution(t *testing.T) {
	list1 := []int{1, 3, 5, 6}
	target1 := 5
	target2 := 2
	target3 := 7
	target4 := 0

	if searchInsert(list1, target1) != 2 {
		t.Error("test error")
	}
	if searchInsert(list1, target2) != 1 {
		t.Error("test error")
	}
	if searchInsert(list1, target3) != 4 {
		t.Error("test error")
	}
	if searchInsert(list1, target4) != 0 {
		t.Error("test error")
	}

	list2 := []int{}
	target1 = 1
	if searchInsert(list2, target1) != 0 {
		t.Error("test error")
	}

	list3 := []int{1}
	target1 = 0
	if searchInsert(list3, target1) != 0 {
		t.Error("test error")
	}

}
