package issue_0114

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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)

	left, right := root.Left, root.Right
	root.Left, root.Right = nil, left

	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = right
}
