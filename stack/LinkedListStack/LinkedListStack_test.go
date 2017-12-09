package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListStack(t *testing.T) {
	assert := assert.New(t)

	stack := NewLinkedListStack()
	assert.Equal(uint(0), stack.data.Size())
}

func TestPush(t *testing.T) {
	assert := assert.New(t)

	stack := NewLinkedListStack()

	stack.Push(1) // 1
	val, err := stack.data.Get(0)
	assert.Equal(1, val)
	assert.Nil(err)

	stack.Push(2)
	stack.Push(3) // 3->2->1
	val, err = stack.data.Get(0)
	assert.Equal(3, val)
	assert.Nil(err)
	val, err = stack.data.Get(1)
	assert.Equal(2, val)
	assert.Nil(err)
	val, err = stack.data.Get(2)
	assert.Equal(1, val)
	assert.Nil(err)
}

func TestPop(t *testing.T) {
	assert := assert.New(t)

	stack := NewLinkedListStack()

	val, err := stack.Pop()
	assert.Nil(val)
	assert.Equal("Stack is empty", err.Error())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3) // [1 2 3]

	val, err = stack.Pop()
	assert.Equal(3, val)
	assert.Nil(err)
	assert.Equal(uint(2), stack.Size())

	stack.Push(5) // [1 2 5]
	val, err = stack.Pop()
	assert.Equal(5, val)
	assert.Nil(err)
	assert.Equal(uint(2), stack.Size())
}

func TestTop(t *testing.T) {
	assert := assert.New(t)

	stack := NewLinkedListStack()

	val, err := stack.Top()
	assert.Nil(val)
	assert.Equal("Stack is empty", err.Error())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3) // [1 2 3]

	val, err = stack.Top()
	assert.Equal(3, val)
	assert.Nil(err)
	assert.Equal(uint(3), stack.Size())

	stack.Push(5) // [1 2 3 5]
	val, err = stack.Top()
	assert.Equal(5, val)
	assert.Nil(err)
	assert.Equal(uint(4), stack.Size())
}

func TestIsEmpty(t *testing.T) {
	assert := assert.New(t)

	stack := NewLinkedListStack()
	assert.Equal(true, stack.IsEmpty())

	stack.Push(1)
	assert.Equal(false, stack.IsEmpty())
}

func TestSize(t *testing.T) {
	assert := assert.New(t)

	stack := NewLinkedListStack()
	assert.Equal(uint(0), stack.Size())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3) // [1 2 3]
	assert.Equal(uint(3), stack.Size())
}
