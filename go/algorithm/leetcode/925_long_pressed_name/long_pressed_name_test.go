package leetcode925

import "testing"

func TestSolution(t *testing.T) {
	if !isLongPressedName("alex", "aaleex") {
		t.Error("test error")
	}
	if isLongPressedName("saeed", "ssaaedd") {
		t.Error("test error")
	}
	if !isLongPressedName("leelee", "lleeelee") {
		t.Error("test error")
	}
	if !isLongPressedName("laiden", "laiden") {
		t.Error("test error")
	}
	if !isLongPressedName("vtkgn", "vttkgnn") {
		t.Error("test error")
	}
	if !isLongPressedName("vtkgn", "vttkgnn") {
		t.Error("test error")
	}
	if isLongPressedName("dfuyalc", "fuuyallc") {
		t.Error("test error")
	}
	if isLongPressedName("pyplrz", "ppyypllr") {
		t.Error("test error")
	}
}
