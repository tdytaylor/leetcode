package stack

import (
	"container/list"
)

type Stack interface {
	Push(value interface{})
	Pop() (value interface{})
	Peek() (value interface{})
	Len() int
	Empty() bool
}

// 使用container/list实现Stack
type stackDefaultImpl struct {
	data *list.List
}

func (stack *stackDefaultImpl) Push(value interface{}) {
	stack.data.PushBack(value)
}

func (stack *stackDefaultImpl) Pop() (value interface{}) {
	e := stack.data.Back()
	if e != nil {
		stack.data.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *stackDefaultImpl) Peek() (value interface{}) {
	e := stack.data.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (stack *stackDefaultImpl) Len() int {
	return stack.data.Len()
}

func (stack *stackDefaultImpl) Empty() bool {
	return stack.data.Len() == 0
}

func NewStack() Stack {
	return &stackDefaultImpl{
		data: list.New(),
	}
}
