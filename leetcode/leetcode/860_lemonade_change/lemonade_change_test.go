package leetcode860

import "testing"

func TestSolution(t *testing.T) {
	if lemonadeChange([]int{5, 5, 5, 10, 20}) != true {
		t.Error("test error")
	}
	if lemonadeChange([]int{5, 5, 10}) != true {
		t.Error("test error")
	}
	if lemonadeChange([]int{10, 10}) != false {
		t.Error("test error")
	}
	if lemonadeChange([]int{5, 5, 10, 10, 20}) != false {
		t.Error("test error")
	}
}
