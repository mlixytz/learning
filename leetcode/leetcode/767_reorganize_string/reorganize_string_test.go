package leetcode767

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if reorganizeString("aab") != "aba" {
		t.Error("test error")
	}
	if reorganizeString("aaab") != "" {
		t.Error("test error")
	}
	if reorganizeString("baaba") != "ababa" {
		t.Error("test error")
	}
	if reorganizeString("baabb") != "babab" {
		t.Error("test error")
	}
	if reorganizeString("vvvlo") != "vlvov" {
		t.Error("test error")
	}

}
