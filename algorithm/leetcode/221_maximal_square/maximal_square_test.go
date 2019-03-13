package leetcode221

import (
	"testing"
)

func TestSolution(t *testing.T) {
	if maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}) != 4 {
		t.Error("test error")
	}
	if maximalSquare([][]byte{{'0', '0', '0'}, {'0', '0', '0'}, {'1', '1', '1'}}) != 1 {
		t.Error("test error")
	}
	if maximalSquare([][]byte{{'1', '1', '1', '1'}, {'1', '1', '1', '1'}, {'1', '1', '1', '1'}}) != 9 {
		t.Error("test error")
	}
	if maximalSquare([][]byte{{'0', '0', '0', '0', '0'}, {'1', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}) != 1 {
		t.Error("test error")
	}
}
