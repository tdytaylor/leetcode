package issue_0116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

// 头都晕了
func connectTwoNode(left *Node, right *Node) {
	if right == nil || left == nil {
		return
	}
	left.Next = right
	connectTwoNode(left.Left, left.Right)
	connectTwoNode(right.Left, right.Right)
	connectTwoNode(left.Right, right.Left)
}
