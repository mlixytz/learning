package leetcode692

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if !isEqual(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2), []string{"i", "love"}) {
		t.Error("test error")
	}
	if !isEqual(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4), []string{"the", "is", "sunny", "day"}) {
		t.Error("test error")
	}
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
