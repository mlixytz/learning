package leetcode187

import "testing"

func TestSolution(t *testing.T) {
	if !isEqual(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"), []string{"AAAAACCCCC", "CCCCCAAAAA"}) {
		t.Error("test error")
	}
	if !isEqual(findRepeatedDnaSequences("AAAAAAAAAAAA"), []string{"AAAAAAAAAA"}) {
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
