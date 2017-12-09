package linkedList

import (
	"fmt"
)

type LinkedNode struct {
	data interface{}
	next *LinkedNode
}

// New creates and return a new empty LinkedNode
func NewLinkedNode() *LinkedNode {
	return &LinkedNode{nil, nil}
}

// New creates and return the pointer to the created LinkedNode with the provided data
func NewLinkedNodeWithVal(data interface{}) *LinkedNode {
	return &LinkedNode{data, nil}
}

// Val returns the value of the LinkedNode
func (node *LinkedNode) Val() interface{} {
	return node.data
}

// SetVal set the provided value of the LinkedNode
func (node *LinkedNode) SetVal(val interface{}) {
	node.data = val
}

// Next returns the pointer to the LinkedNode pointed by the next pointer,
// nil if the pointer has not assigned any value
func (node *LinkedNode) Next() *LinkedNode {
	return node.next
}

// SetNext updates the next pointer to the specified LinkedNode
func (node *LinkedNode) SetNext(next *LinkedNode) {
	node.next = next
}

// InsertAfter insert a LinkedNode with the specified data after the current
// node and return the pointer to the LinkedNode just inserted
func (node *LinkedNode) InsertAfter(data interface{}) *LinkedNode {
	tmp := node.next
	node.next = NewLinkedNodeWithVal(data)
	node.next.next = tmp

	return node.next
}

func (node *LinkedNode) String() string {
	return fmt.Sprint(node.data)
}
