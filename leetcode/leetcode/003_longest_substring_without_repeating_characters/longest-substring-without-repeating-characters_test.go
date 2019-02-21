package leetcode003

import (
	"testing"
)

func TestSolution(t *testing.T) {
	s := "tmmzuxt"
	result := lengthOfLongestSubstring(s)
	if result != 5 {
		t.Error("test Error")
	}
	resultBs := lengthOfLongestSubstringBestSolution(s)
	if resultBs != 5 {
		t.Error("test best Solution Error")
	}
}
