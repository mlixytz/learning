package leetcode543

import (
	"fmt"
	"testing"

	"github.com/mlixytz/LeetCode-go/utils"
)

func TestSolution(t *testing.T) {
	if diameterOfBinaryTree(utils.CreateTreeByArray([]int{1, 2, 3, 4, 5})) != 3 {
		fmt.Println(diameterOfBinaryTree(utils.CreateTreeByArray([]int{1, 2, 3, 4, 5})))
		t.Error("test error")
	}
}
