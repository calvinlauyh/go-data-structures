package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedNode(t *testing.T) {
	assert := assert.New(t)

	node := NewLinkedNode()
	assert.Nil(node.data)
	assert.Nil(node.next)
}

func TestNewLinkedNodeWithVal(t *testing.T) {
	assert := assert.New(t)

	node := NewLinkedNodeWithVal(1)
	assert.Equal(1, node.data)
	assert.Nil(node.next)
}

func TestLinkedNodeVal(t *testing.T) {
	assert := assert.New(t)

	node := NewLinkedNodeWithVal(1)
	assert.Equal(1, node.Val())
}

func TestLinkedNodeSetVal(t *testing.T) {
	assert := assert.New(t)

	node := NewLinkedNodeWithVal(1)
	node.SetVal(2)
	assert.Equal(2, node.data)
}

func TestLinkedNodeSetNext(t *testing.T) {
	assert := assert.New(t)

	node := NewLinkedNodeWithVal(1)
	node2 := NewLinkedNodeWithVal(2)
	node.SetNext(node2)
	assert.Equal(node2, node.next)
}

func TestLinkedNodeNext(t *testing.T) {
	assert := assert.New(t)

	node := NewLinkedNodeWithVal(1)
	node2 := NewLinkedNodeWithVal(2)
	node.SetNext(node2)
	assert.Equal(node2, node.Next())
}

func TestLinkedNodeInsertAfter(t *testing.T) {
	assert := assert.New(t)

	node := NewLinkedNodeWithVal(1)
	node3 := node.InsertAfter(3)
	assert.Equal(3, node.next.data)
	assert.Equal(node3, node.next)
	assert.Nil(node.next.next)

	node2 := node.InsertAfter(2)
	assert.Equal(2, node.next.data)
	assert.Equal(node2, node.next)
	assert.Equal(node3, node.next.next)
	assert.Nil(node.next.next.next)
}

type dummy struct{}

func (d dummy) String() string {
	return "dummy"
}
func TestLinkedNodeString(t *testing.T) {
	assert := assert.New(t)

	node := NewLinkedNodeWithVal(1)
	assert.Equal("1", node.String())

	node = NewLinkedNodeWithVal(dummy{})
	assert.Equal("dummy", node.String())
}
