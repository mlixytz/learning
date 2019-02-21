package leettcode682

import "testing"

func TestSolution(t *testing.T) {
	if calPoints([]string{"5", "2", "C", "D", "+"}) != 30 {
		t.Error("test error")
	}
	if calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}) != 27 {
		t.Error("test error")
	}
}
