package issue_0092

import (
	. "github.com/tdytaylor/leetcode/structure"
)

var (
	successor *ListNode
)

// 递归解决
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	lastNode := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return lastNode
}
