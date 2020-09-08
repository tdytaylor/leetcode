package stack

type Stack interface {
	Push(value interface{})
	Pop(value interface{})
	Peek(value interface{})
}
