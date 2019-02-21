package leetcode889

/*
   [root][......left......][...right..]  ---pre
   [......left......][...right..][root]  ---post
*/

import (
	"github.com/mlixytz/LeetCode-go/utils"
)

type TreeNode = utils.TreeNode

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	if len(pre) == 1 {
		return root
	}
	left := pre[1]
	i := 0
	for ; i < len(post); i++ {
		if post[i] == left {
			break
		}
	}
	root.Left = constructFromPrePost(pre[1:1+i+1], post[:i+1])
	root.Right = constructFromPrePost(pre[2+i:], post[i+1:len(post)-1])
	return root
}
