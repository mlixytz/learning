package leetcode012

import "testing"

func TestSolution(t *testing.T) {
	if intToRoman(3) != "III" {
		t.Error("test error")
	}
	if intToRoman(58) != "LVIII" {
		t.Error("test error")
	}
	if intToRoman(1994) != "MCMXCIV" {
		t.Error("test error")
	}
}
