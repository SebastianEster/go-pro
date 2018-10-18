package main

type Stack struct {
	data []interface{}
}

func newStack() *Stack {
	return new(Stack)
}

func (stack *Stack) push(object interface{}) {
	stack.data = append(stack.data, object)
}

func (stack *Stack) pop() interface{} {
	if len(stack.data) == 0 {
		panic("Cannot pop from an empty stack!")
	}

	ret := stack.data[len(stack.data) - 1]
	stack.data = stack.data[:len(stack.data) - 1]
	return ret
}

func (stack *Stack) getAt(index int) interface{} {
	return stack.data[index]
}

func (stack *Stack) size() int {
	return len(stack.data)
}