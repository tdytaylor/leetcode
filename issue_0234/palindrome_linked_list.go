package issue_0234

import (
	. "github.com/tdytaylor/leetcode/structure"
)

var (
	left *ListNode
)

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}

// 精辟
func traverse(head *ListNode) bool {
	if head == nil {
		return true
	}
	b := traverse(head.Next)
	b = b && (head.Val == left.Val)
	left = left.Next
	return b
}
