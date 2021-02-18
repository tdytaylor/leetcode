package main

import (
	"fmt"
	. "github.com/tdytaylor/leetcode/structure"
)

func main() {
	l1 := &ListNode{
		Val: 2,
	}
	l1.Append(4).Append(3).Append(6).Append(7)
	fmt.Println(l1)
	// fmt.Println(reverseList3(l1))
	fmt.Println(reverseList(l1))
}

// LeetCode最优解
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, next, isHead := head, head.Next, true
	for next != nil {
		if isHead {
			cur.Next = nil
			isHead = false
		}
		next.Next, cur, next = cur, next, next.Next
	}
	return cur
}

// 时间复杂度为 O(n) 空间复杂度为 O(1)
func reverseList2(head *ListNode) *ListNode {
	// 把链表拆分成两段，然后遍历第二段，依次把第二段的节点放到第一个链表的头结点
	var curr, next *ListNode
	// 如果链表为空或者只有一个元素，直接返回
	if head == nil || head.Next == nil {
		return head
	}
	curr = head
	next = curr.Next
	curr.Next = nil
	for next != nil {
		temp := next
		next = next.Next
		temp.Next = curr
		curr = temp
	}
	return curr
}

// 这种方式不会断链
func reverseList3(head *ListNode) *ListNode {
	var curr, next *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	curr = head
	for curr.Next != nil {
		next = curr.Next
		curr.Next = next.Next
		next.Next = head
		head = next
	}
	return head
}

//
func reverseListByRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	recursion := reverseListByRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return recursion
}
