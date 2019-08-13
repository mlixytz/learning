package leetcode762

import "testing"

func TestSolution(t *testing.T) {
	if countPrimeSetBits(6, 10) != 4 {
		t.Error("test error")
	}
	if countPrimeSetBits(10, 15) != 5 {
		t.Error("test error")
	}
}
