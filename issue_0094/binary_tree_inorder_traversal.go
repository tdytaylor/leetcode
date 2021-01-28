package issue_0094

import (
	. "github.com/tdytaylor/leetcode/structure"
	"github.com/tdytaylor/leetcode/structure/stack"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归遍历
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

// 树前序遍历-Stack
func inorderTraversalByStack(root *TreeNode) []int {
	var res []int
	stack := stack.NewStack()
	node := root
	for node != nil || !stack.Empty() {
		// 迭代访问节点的左孩子，并入栈
		for node != nil {
			stack.Push(node)
			node = node.Left
		}

		if !stack.Empty() {
			node = stack.Pop().(*TreeNode)

			node = node.Right
		}
	}
	return res
}
