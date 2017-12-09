package stack

import (
	"errors"
)

const (
	ARRAYSTACK_DEFAULT_CAP = 30
)

type ArrayStack struct {
	data []interface{}
	top  int
}

// Create and return a new Array Stack with the default capacity of the
// underlying array.
func NewArrayStack() *ArrayStack {
	return NewArrayStackWithDefaultCap(ARRAYSTACK_DEFAULT_CAP)
}

// Create and return a new Array Stack with specified default capacity of the
// underlying array. The default capacity is just for initialization and the
// array will expands itself when the number of elements in stack exceed the
// size
func NewArrayStackWithDefaultCap(defaultCap uint) *ArrayStack {
	data := make([]interface{}, 0, defaultCap)
	return &ArrayStack{data, -1}
}

// Push inserts an element to the top of stack
func (stack *ArrayStack) Push(val interface{}) {
	stack.data = append(stack.data, val)
	stack.top++
}

// Pop removes and returns the topmost element from stack, error if the stack
// is empty
func (stack *ArrayStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	data := stack.data[stack.top]
	stack.data = stack.data[:stack.top]
	stack.top--
	return data, nil
}

// Top returns the topmost element from stack, error if the stack is empty
func (stack *ArrayStack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	return stack.data[stack.top], nil
}

// IsEmpty returns whether the stack is empty
func (stack *ArrayStack) IsEmpty() bool {
	return stack.top == -1
}

// Size returns the number of elements in the stack
func (stack *ArrayStack) Size() uint {
	return uint(stack.top + 1)
}
