package josephuspermutation

import "testing"

func TestJosephusPermutation1(t *testing.T) {
	if josephusPermutation1(2, 1, 1) != 1 {
		t.Error("test error")
	}
	if josephusPermutation1(3, 1, 1) != 3 {
		t.Error("test error")
	}
	if josephusPermutation1(3, 1, 2) != 2 {
		t.Error("test error")
	}
	if josephusPermutation1(4, 3, 2) != 3 {
		t.Error("test error")
	}

	if josephusPermutation2(2, 1) != 1 {
		t.Error("test error")
	}
	if josephusPermutation2(3, 1) != 3 {
		t.Error("test error")
	}
	if josephusPermutation2(3, 2) != 2 {
		t.Error("test error")
	}
}
