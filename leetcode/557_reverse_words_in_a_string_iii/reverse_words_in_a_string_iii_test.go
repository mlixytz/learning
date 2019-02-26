package leetcode557

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	if reverseWords("Let's take LeetCode contest") != "s'teL ekat edoCteeL tsetnoc" {
		fmt.Println(reverseWords("Let's take LeetCode contest"))
		t.Error("test error")
	}
}
