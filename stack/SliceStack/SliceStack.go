package stack

import (
	"errors"
)

const (
	SLICESTACK_DEFAULT_CAP = 30
)

type SliceStack struct {
	data []interface{}
	top  int
}

// Create and return a new Array Stack with the default capacity of the
// underlying array.
func NewSliceStack() *SliceStack {
	return NewSliceStackWithDefaultCap(SLICESTACK_DEFAULT_CAP)
}

// Create and return a new Array Stack with specified default capacity of the
// underlying array. The default capacity is just for initialization and the
// array will expands itself when the number of elements in stack exceed the
// size
func NewSliceStackWithDefaultCap(defaultCap uint) *SliceStack {
	data := make([]interface{}, 0, defaultCap)
	return &SliceStack{data, -1}
}

// Push inserts an element to the top of stack
func (stack *SliceStack) Push(val interface{}) {
	stack.data = append(stack.data, val)
	stack.top++
}

// Pop removes and returns the topmost element from stack, error if the stack
// is empty
func (stack *SliceStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	data := stack.data[stack.top]
	stack.data = stack.data[:stack.top]
	stack.top--
	return data, nil
}

// Top returns the topmost element from stack, error if the stack is empty
func (stack *SliceStack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	return stack.data[stack.top], nil
}

// IsEmpty returns whether the stack is empty
func (stack *SliceStack) IsEmpty() bool {
	return stack.top == -1
}

// Size returns the number of elements in the stack
func (stack *SliceStack) Size() uint {
	return uint(stack.top + 1)
}
