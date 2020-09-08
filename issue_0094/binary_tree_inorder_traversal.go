package issue_0094

import (
	"github.com/emirpasic/gods/stacks/arraystack"
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

// 树遍历-stack
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	cur := root
	stack := arraystack.New()
	for cur != nil || !stack.Empty() {
		for cur.Left != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		node, _ := stack.Pop()
	}
}
