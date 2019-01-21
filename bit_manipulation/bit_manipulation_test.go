package bitmanipulation

import "testing"

func TestSolution(t *testing.T) {
	if countOne(10) != 2 || countOne(0) != 0 {
		t.Error("test error")
	}

	if isPowerOfFour(10) || !isPowerOfFour(64) {
		t.Error("test error")
	}

	if largestPower(10) != 8 || largestPower(6000) != 4096 {
		t.Error("test error")
	}

}
