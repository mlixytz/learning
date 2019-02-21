package leetcode763

import (
	"fmt"
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if !utils.IsArrayEqual(partitionLabels("ababcbacadefegdehijhklij"), []int{9, 7, 8}) {
		fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
		t.Error("test error")
	}
}
