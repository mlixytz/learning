package leetcode242

import "testing"

func TestSolution(t *testing.T) {
	if !isAnagram("anagram", "nagaram") {
		t.Error("test error")
	}
	if isAnagram("rat", "car") {
		t.Error("test error")
	}
	if !isAnagram("", "") {
		t.Error("test error")
	}
}
