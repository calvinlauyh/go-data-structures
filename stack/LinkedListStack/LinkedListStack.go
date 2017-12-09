package stack

import (
	"errors"
	. "github.com/yuhlau/go-data-structures/linkedList"
)

type LinkedListStack struct {
	data *LinkedList
}

// Create and return a new ListedList Stack
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{NewLinkedList()}
}

// Push inserts an element to the top of stack
func (stack *LinkedListStack) Push(val interface{}) {
	stack.data.Insert(0, val)
}

// Pop removes and returns the topmost element from stack, error if the stack
// is empty
func (stack *LinkedListStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	val, err := stack.data.Delete(0)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// Top returns the topmost element from stack, error if the stack is empty
func (stack *LinkedListStack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	val, err := stack.data.Get(0)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// IsEmpty returns whether the stack is empty
func (stack *LinkedListStack) Size() uint {
	return stack.data.Size()
}

// Size returns the number of elements in the stack
func (stack *LinkedListStack) IsEmpty() bool {
	return stack.data.Size() == 0
}
