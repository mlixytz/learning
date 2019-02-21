package leetcode451

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if frequencySort("tree") != "eert" && frequencySort("tree") != "eetr" {
		t.Error("test error")
	}
}
