package issue_0226

import (
	. "github.com/tdytaylor/leetcode/structure"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归实现
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	root.Left = root.Right
	root.Right = left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
