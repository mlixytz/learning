package leetcode190

import "testing"

func TestSolution(t *testing.T) {
	if reverseBits(43261596) != 964176192 {
		t.Error("test errror")
	}
	if reverseBits(4294967293) != 3221225471 {
		t.Error("test errror")
	}
}
