package leetcode933

import "testing"

func TestSolution(t *testing.T) {
	obj := Constructor()
	if obj.Ping(1) != 1 {
		t.Error("test error")
	}
	if obj.Ping(100) != 2 {
		t.Error("test error")
	}
	if obj.Ping(3001) != 3 {
		t.Error("test error")
	}
	if obj.Ping(3002) != 3 {
		t.Error("test error")
	}

}
