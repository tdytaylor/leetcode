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

// 精辟：学习思想。时间和空间复杂度都是：O(N)
func traverse(head *ListNode) bool {
	if head == nil {
		return true
	}
	b := traverse(head.Next)
	b = b && (head.Val == left.Val)
	left = left.Next
	return b
}

// 第二种方式
func isPalindrome2(head *ListNode) bool {
	slow, fast, left := head, head, head
	var slowPre *ListNode
	for fast != nil && fast.Next != nil {
		slowPre = slow // 恢复链表
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	right := reverse(slow)
	rightHead := right // 恢复链表

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}
	slowPre.Next = reverse(rightHead) // 恢复链表
	return true
}

func reverse(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nex := cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}
