package leetcode703

import "testing"

func TestSolution(t *testing.T) {
	obj1 := Constructor(1, []int{})
	if obj1.Add(3) != 3 {
		t.Error("test error")
	}
	obj := Constructor(3, []int{4, 5, 8, 2})
	if obj.Add(3) != 4 {
		t.Error("test error")
	}
	if obj.Add(5) != 5 {
		t.Error("test error")
	}
	if obj.Add(10) != 5 {
		t.Error("test error")
	}
	if obj.Add(9) != 8 {
		t.Error("test error")
	}
	if obj.Add(4) != 8 {
		t.Error("test error")
	}
	obj2 := Constructor(2, []int{0})
	if obj2.Add(-1) != -1 {
		t.Error("test error")
	}
	if obj2.Add(1) != 0 {
		t.Error("test error")
	}
	if obj2.Add(-2) != 0 {
		t.Error("test error")
	}
	if obj2.Add(-4) != 0 {
		t.Error("test error")
	}
	if obj2.Add(3) != 1 {
		t.Error("test error")
	}

}
