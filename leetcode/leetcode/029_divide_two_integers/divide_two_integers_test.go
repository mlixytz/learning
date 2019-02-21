package leetcode029

import "testing"

func TestSolution(t *testing.T) {
	if divide(9, 3) != 3 {
		t.Error("test error")
	}
	if divide(10, 3) != 3 {
		t.Error("test error")
	}
	if divide(3, 10) != 0 {
		t.Error("test error")
	}
	if divide(10, -3) != -3 {
		t.Error("test error")
	}
	if divide(-10, 3) != -3 {
		t.Error("test error")
	}
	if divide(-2147483648, -1) != 2147483647 {
		t.Error("test error")
	}

	if divideBestSolution(9, 3) != 3 {
		t.Error("test error")
	}
	if divideBestSolution(10, 3) != 3 {
		t.Error("test error")
	}
	if divideBestSolution(3, 10) != 0 {
		t.Error("test error")
	}
	if divideBestSolution(10, -3) != -3 {
		t.Error("test error")
	}
	if divideBestSolution(-10, 3) != -3 {
		t.Error("test error")
	}
	if divideBestSolution(-2147483648, -1) != 2147483647 {
		t.Error("test error")
	}
	if divideBestSolution(1004958205, -2137325331) != 1004958205/-2137325331 {
		t.Error("test error")
	}

}
