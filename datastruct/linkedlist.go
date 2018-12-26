package datastruct

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// append the node to linkedlist
func (list *ListNode) Append(value int) *ListNode {
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
