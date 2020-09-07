package issue_0110

import (
	. "github.com/tdytaylor/leetcode/structure"
	"math"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	return math.Abs(float64(leftHeight-rightHeight)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Right), depth(root.Left)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
