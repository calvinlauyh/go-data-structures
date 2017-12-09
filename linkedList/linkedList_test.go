package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()

	assert.Equal(NewLinkedNode(), list.head)
	assert.Nil(list.head.Next())
}

func TestLinkedListHead(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()

	val, ok := list.Head()
	assert.Nil(val)
	assert.Equal(false, ok)

	list.Append(1)
	val, ok = list.Head()
	assert.Equal(1, val)
	assert.Equal(true, ok)
}

func TestLinkedListTail(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()

	val, ok := list.Tail()
	assert.Nil(val)
	assert.Equal(false, ok)

	list.Append(1)
	list.Append(2) // [1 2]
	val, ok = list.Tail()
	assert.Equal(2, val)
	assert.Equal(true, ok)
}

func TestLinkedListIsEmpty(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	assert.Equal(true, list.IsEmpty())
}

func TestLinkedListSize(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	assert.Equal(uint(0), list.Size())

	list.Append(1) // [1]
	assert.Equal(uint(1), list.Size())

	list.Append(2) // [1 2]
	assert.Equal(uint(2), list.Size())
}

func TestLinkedListGet(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	list.Append(1)
	list.Append(2) // [1 2]

	val, err := list.Get(0)
	assert.Equal(1, val)
	assert.Equal(nil, err)

	val, err = list.Get(1)
	assert.Equal(2, val)
	assert.Equal(nil, err)

	val, err = list.Get(9999)
	assert.Nil(val)
	assert.Equal("Invalid position", err.Error())
}

func TestLinkedListGetNode(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	list.Append(1)
	list.Append(2) // [1 2]

	node, err := list.GetNode(0)
	assert.Equal(1, node.Val())

	node2, err := list.GetNode(1)
	assert.Equal(2, node2.Val())
	assert.Equal(node2, node.Next())

	node, err = list.GetNode(9999)
	assert.Nil(node)
	assert.Equal("Invalid position", err.Error())
}

func TestLinkedListAppend(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	node := list.Append(1) // [1]

	assert.Equal(node, list.head.Next())
	assert.Equal(1, list.head.Next().Val())

	node = list.Append(2) // [1 2]
	assert.Equal(node, list.head.Next().Next())
	assert.Equal(1, list.head.Next().Val())
	assert.Equal(2, list.head.Next().Next().Val())
}

func TestLinkedListInsert(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	node, err := list.Insert(0, 1) // [1]
	assert.Equal(1, list.head.Next().Val())
	assert.Equal(list.head.Next(), node)
	assert.Nil(err)

	node, err = list.Insert(1, 3) // [1 3]
	assert.Equal(1, list.head.Next().Val())
	assert.Equal(3, list.head.Next().Next().Val())
	assert.Equal(list.head.Next().Next(), node)
	assert.Nil(err)

	node, err = list.Insert(1, 2) // [1 2 3]
	assert.Equal(1, list.head.Next().Val())
	assert.Equal(2, list.head.Next().Next().Val())
	assert.Equal(3, list.head.Next().Next().Next().Val())
	assert.Equal(list.head.Next().Next(), node)
	assert.Nil(err)

	node, err = list.Insert(9999, 1)
	assert.Nil(node)
	assert.Equal("Invalid position", err.Error())

}

func TestLinkedListDelete(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)              // [1 2 3]
	node, err := list.Delete(1) // [1 3]

	assert.Equal(2, node)
	assert.Equal(1, list.head.Next().Val())
	assert.Equal(3, list.head.Next().Next().Val())

	node, err = list.Delete(9999)
	assert.Nil(node)
	assert.Equal("Invalid position", err.Error())
}

func TestLinkedListFind(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(1)
	list.Append(1)
	list.Append(2)
	list.Append(3) // [1 2 3 1 1 2 3]

	find9999 := func(val interface{}) bool { return val == 9999 }
	find1 := func(val interface{}) bool { return val == 1 }
	find2 := func(val interface{}) bool { return val == 2 }
	find3 := func(val interface{}) bool { return val == 3 }

	assert.Equal(-1, list.Find(find9999))
	assert.Equal(0, list.Find(find1))
	assert.Equal(1, list.Find(find2))
	assert.Equal(2, list.Find(find3))
}

func TestLinkedListFindOccurence(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(1)
	list.Append(1)
	list.Append(2)
	list.Append(3) // [1 2 3 1 1 2 3]

	find9999 := func(val interface{}) bool { return val == 9999 }
	find1 := func(val interface{}) bool { return val == 1 }

	assert.Equal(-1, list.FindByOccurence(find9999, 1))
	assert.Equal(0, list.FindByOccurence(find1, 1))
	assert.Equal(3, list.FindByOccurence(find1, 2))
	assert.Equal(4, list.FindByOccurence(find1, -1))
	assert.Equal(-1, list.FindByOccurence(find1, 5))
}

func TestLinkedListString(t *testing.T) {
	assert := assert.New(t)

	list := NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	assert.Equal("[1 2 3 4 5]", list.String())
}
