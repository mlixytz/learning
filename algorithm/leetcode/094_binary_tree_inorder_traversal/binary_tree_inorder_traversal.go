package leetcode094

/*
迭代方式实现中序遍历：
	使用stack保存节点, cur 代表当前节点, ans代表结果数组
	1. 节点如果包含左节点，将左节点添加到stack中，并使cur指向curl.Left 直到左节点遍历完成。
	2. pop stack 的节点作为 cur。 并把cur 添加到ans中， 然后将cur指向cur.Right。重复步骤1
	3. 直到stack为空时ans便是最终结果
*/

import (
	"container/list"

	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := list.New()
	cur := root
	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushFront(cur)
			cur = cur.Left
		}
		cur = stack.Remove(stack.Front()).(*TreeNode)
		ans = append(ans, cur.Val)
		cur = cur.Right
	}
	return ans
}
