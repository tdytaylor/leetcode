package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{
		Val: 2,
	}
	l1.append(4).append(3)
	l2 := &ListNode{
		Val: 5,
	}
	l2.append(6).append(4).append(3)
	//fmt.Println(l1)
	// fmt.Println(l2)
	fmt.Println(addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
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

func (list *ListNode) append(value int) *ListNode {
	node := &ListNode{
		Val: value,
	}
	list.Next = node
	return node
}

func (list *ListNode) String() string {
	result := list.print()
	nextNode := list
	for nextNode.Next != nil {
		result += nextNode.Next.print()
		nextNode = nextNode.Next
	}
	return result
}

func (list *ListNode) print() string {
	return fmt.Sprintf(" %d ", list.Val)

}

// convert the ListNode to int slice
func (list *ListNode) convertToSlice() []int {
	arr := make([]int, 5, 128)
	i := 0
	arr[i] = list.Val
	nextNode := list
	for nextNode.Next != nil {
		i += 1
		arr[i] = nextNode.Next.Val
		nextNode = nextNode.Next
	}
	return arr[0 : i+1]
}

func reverse(slice []int) *[]int {
	for i := 0; i < int(len(slice)/2); i++ {
		slice[i], slice[len(slice)-(i+1)] = slice[len(slice)-(i+1)], slice[i]
	}
	return &slice
}
