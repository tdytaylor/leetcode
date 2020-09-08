package main

import (
	"fmt"
	. "github.com/tdytaylor/leetcode/structure"
)

func main() {
	l1 := &ListNode{
		Val: 2,
	}
	l1.Append(4).Append(3)
	l2 := &ListNode{
		Val: 5,
	}
	l2.Append(6).Append(4).Append(3)
	//fmt.Println(l1)
	// fmt.Println(l2)
	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 哨兵
	head := &ListNode{
		Val: 0,
	}
	carry := 0
	list1, list2, curr := l1, l2, head
	for list1 != nil || list2 != nil {
		x, y := 0, 0
		if list1 != nil {
			x = list1.Val
		}
		if list2 != nil {
			y = list2.Val
		}
		sum := x + y + carry
		// 处理进位
		carry = int(sum / 10)

		value := sum % 10
		curr.Next = &ListNode{
			Val: value,
		}
		curr = curr.Next
		if list1 != nil {
			list1 = list1.Next
		}
		if list2 != nil {
			list2 = list2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}
