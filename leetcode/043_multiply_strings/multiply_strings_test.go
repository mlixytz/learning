package leetcode043

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if multiply("2", "3") != "6" {
		t.Error("test error")
	}
	if multiply("123", "456") != "56088" {
		fmt.Println(multiply("123", "456"))
		t.Error("test error")
	}
}
