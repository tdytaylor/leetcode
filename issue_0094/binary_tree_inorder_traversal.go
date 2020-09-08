package issue_0094

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
func inorderTraversal(root *TreeNode) []int {
	var res []int
	return traversal(res, root)
}

func traversal(slice []int, root *TreeNode) (rightSlice []int) {
	if root != nil {
		slice = traversal(slice, root.Left)
		slice = append(slice, root.Val)
		slice = traversal(slice, root.Right)
	}
	rightSlice = slice
	return
}
