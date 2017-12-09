package linkedList

import (
	"bytes"
	"errors"
	"strings"
)

type LinkedList struct {
	head *LinkedNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewLinkedNode()}
}

// Head returns the first element of the list, second returned value will be
// false if the list is empty
func (list *LinkedList) Head() (interface{}, bool) {
	if list.head.Next() == nil {
		return nil, false
	}
	return list.head.Next().Val(), true
}

// Tail returns the last element of the list, second returned value will be
// false if the list is empty
func (list *LinkedList) Tail() (interface{}, bool) {
	if list.head.Next() == nil {
		return nil, false
	}
	tail := list.head.Next()
	for tail.Next() != nil {
		tail = tail.Next()
	}
	return tail.Val(), true
}

// IsEmpty returns whether the list is empty
func (list *LinkedList) IsEmpty() bool {
	return list.head.Next() == nil
}

// Size returns the number of element(s) in the list
func (list *LinkedList) Size() uint {
	var count uint = 0
	current := list.head.Next()
	for current != nil {
		current = current.Next()
		count++
	}
	return count
}

// Get returns the element at the specified position, or error if the position
// is invalid
func (list *LinkedList) Get(pos uint) (interface{}, error) {
	current := list.head.Next()
	for i := uint(0); i < pos && current != nil; i++ {
		current = current.Next()
	}
	if current == nil {
		return nil, errors.New("Invalid position")
	}
	return current.Val(), nil
}

// GetNode returns the pointer to the LinkedNode at the spcified position, or
// error if the position is invalid
func (list *LinkedList) GetNode(pos uint) (*LinkedNode, error) {
	current := list.head.Next()
	for i := uint(0); i < pos && current != nil; i++ {
		current = current.Next()
	}
	if current == nil {
		return nil, errors.New("Invalid position")
	}
	return current, nil
}

// Append inserts the provided data to the end of the list and return the
// pointer to the LinkedNode just appended
func (list *LinkedList) Append(data interface{}) *LinkedNode {
	current := list.head
	for current.Next() != nil {
		current = current.Next()
	}
	return current.InsertAfter(data)
}

// Insert inserts the provided data to specified position of the list, and
// return the pointer to LinkedNode, or error if the position is invalid
func (list *LinkedList) Insert(pos uint, data interface{}) (*LinkedNode, error) {
	previous := list.head
	i := uint(0)
	for ; i < pos && previous.Next() != nil; i++ {
		previous = previous.Next()
	}
	if i < pos {
		return nil, errors.New("Invalid position")
	}
	return previous.InsertAfter(data), nil
}

// Delete removes an element at the specified position and returns the deleted
// element, or error if the position is invalid
func (list *LinkedList) Delete(pos uint) (interface{}, error) {
	previous := list.head
	i := uint(0)
	for ; i < pos && previous.Next() != nil; i++ {
		previous = previous.Next()
	}
	if i < pos {
		return nil, errors.New("Invalid position")
	}
	tmp := previous.Next().Val()
	previous.SetNext(previous.Next().Next())

	return tmp, nil
}

// Find searches the LinkedList for the first element satisfying the provided
// function and return the index (starting from 0)
func (list *LinkedList) Find(fn func(interface{}) bool) int {
	current := list.head
	for i := 0; current.Next() != nil; i++ {
		current = current.Next()
		if fn(current.Val()) {
			return i
		}
	}
	return -1
}

func (list *LinkedList) FindByOccurence(fn func(interface{}) bool, occurence int) int {
	allMatches := make([]int, 0, list.Size()/2)
	current := list.head
	j := 0
	for i := 0; current.Next() != nil; i++ {
		current = current.Next()
		if fn(current.Val()) {
			j++
			// match position occurence value
			if j == occurence {
				return i
			}
			allMatches = append(allMatches, i)
		}
	}
	if occurence < 0 {
		// Convert the negative occurence to the corresponding slice index
		normal := len(allMatches) + occurence
		if normal < 0 {
			// If normal is less than 0, it means the last n occurences does not
			// exist
			return -1
		}
		return allMatches[normal]
	}
	return -1
}

func (list *LinkedList) String() string {
	var b bytes.Buffer
	els := make([]string, 0, list.Size())

	b.WriteString("[")
	current := list.head.Next()
	for current != nil {
		els = append(els, current.String())
		current = current.Next()
	}
	b.WriteString(strings.Join(els, " "))
	b.WriteString("]")

	return b.String()
}
