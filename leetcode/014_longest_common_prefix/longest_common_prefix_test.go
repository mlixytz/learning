package leetcode014

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if longestCommonPrefix([]string{"flower", "flow", "flight"}) != "fl" {
		fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
		t.Error("test error")
	}
	if longestCommonPrefix([]string{"dog", "racecar", "car"}) != "" {
		t.Error("test error")
	}
}
