package leetcode744

import "testing"

func TestSolution(t *testing.T) {
	letters := []byte{'c', 'f', 'j'}
	letters2 := []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}
	if (nextGreatestLetter(letters, 'a')) != 'c' {
		t.Error("test error")
	}
	if (nextGreatestLetter(letters, 'c')) != 'f' {
		t.Error("test error")
	}
	if (nextGreatestLetter(letters, 'd')) != 'f' {
		t.Error("test error")
	}
	if (nextGreatestLetter(letters, 'j')) != 'c' {
		t.Error("test error")
	}
	if (nextGreatestLetter(letters, 'k')) != 'c' {
		t.Error("test error")
	}
	if (nextGreatestLetter(letters2, 'e')) != 'n' {
		t.Error("test error")
	}
}
