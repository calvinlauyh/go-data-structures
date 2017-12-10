package linkedListTree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/yuhlau/go-data-structures/tree"
)

func TestNewLinkedListTree(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(1)
	assert.Nil(tree.parent)
	assert.Equal(1, tree.val)
	assert.Nil(tree.firstChild)
}

func TestNewLinkedListTreeNode(t *testing.T) {
	assert := assert.New(t)

	parent := NewLinkedListTree(1)
	tree := NewLinkedListTreeNode(2, parent)
	assert.Equal(parent, tree.parent)
	assert.Equal(2, tree.val)
	assert.Nil(tree.firstChild)
}

func TestSetVal(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(1)
	tree.SetVal(2)
	assert.Equal(2, tree.val)
}

func TestVal(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(1)
	assert.Equal(1, tree.Val())
}

func TestSetParent(t *testing.T) {
	assert := assert.New(t)

	parent := NewLinkedListTree(1)
	tree := NewLinkedListTree(2)
	err := tree.SetParent(parent)
	assert.Equal(parent, tree.parent)
	assert.Nil(err)

	tree = NewLinkedListTreeNode(2, parent)
	tree.AppendSibling(3)
	err = tree.SetParent(nil)
	assert.Equal("Tree cannot be root because it already has sibling", err.Error())
}

func TestParent(t *testing.T) {
	assert := assert.New(t)

	parent := NewLinkedListTree(1)
	tree := NewLinkedListTree(2)
	tree.SetParent(parent)
	assert.Equal(parent, tree.Parent())
}

func TestIsRoot(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(2)
	assert.Equal(true, tree.IsRoot())

	parent := NewLinkedListTree(1)
	tree.SetParent(parent)
	assert.Equal(false, tree.IsRoot())
}

func TestAppendSibling(t *testing.T) {
	assert := assert.New(t)

	parent := NewLinkedListTree(1)
	sibling, err := parent.AppendSibling(3)
	assert.Nil(sibling)
	assert.Equal("Root cannot have sibling", err.Error())

	tree := NewLinkedListTreeNode(2, parent)
	sibling1, err := tree.AppendSibling(3)
	assert.Equal(3, sibling1.Val())
	assert.Equal(parent, sibling1.Parent())
	assert.Equal(sibling1, tree.NextSibling())
	assert.Equal(tree, sibling1.PrevSibling())
	assert.Nil(err)

	sibling2, err := tree.AppendSibling(4)
	assert.Equal(4, sibling2.Val())
	assert.Equal(parent, sibling2.Parent())
	assert.Equal(sibling2, sibling1.NextSibling())
	assert.Equal(sibling1, sibling2.PrevSibling())
	assert.Nil(err)
}

func TestAppendChild(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(1)
	child1 := tree.AppendChild(2)
	assert.Equal(child1, tree.FirstChild())
	assert.Equal(tree, child1.Parent())

	child2 := tree.AppendChild(3)
	assert.Equal(child2, tree.FirstChild().NextSibling())
	assert.Equal(tree, child1.Parent())

	child3 := tree.AppendChild(4)
	assert.Equal(child3, tree.FirstChild().NextSibling().NextSibling())
	assert.Equal(tree, child1.Parent())
}

func TestSetPrevSibling(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(3)
	err := tree.SetPrevSibling(NewLinkedListTreeNode(2, tree))
	assert.Equal("Root cannot have sibling", err.Error())

	parent := NewLinkedListTree(1)
	tree.SetParent(parent)
	err = tree.SetPrevSibling(NewLinkedListTree(2))
	assert.Equal("Cannot set a root to be another tree's sibling", err.Error())

	sibling := NewLinkedListTreeNode(2, tree)
	err = tree.SetPrevSibling(sibling)
	assert.Nil(err)
	assert.Equal(sibling, tree.PrevSibling())

	err = tree.SetPrevSibling(nil)
	assert.Nil(err)
	assert.Nil(tree.PrevSibling())
}

func TestPrevSibling(t *testing.T) {
	assert := assert.New(t)

	parent := NewLinkedListTree(1)
	tree := NewLinkedListTree(2)
	tree.SetParent(parent)
	sibling, _ := tree.AppendSibling(2)
	assert.Equal(tree, sibling.PrevSibling())
}

func TestSetNextSibling(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(2)
	err := tree.SetNextSibling(NewLinkedListTreeNode(3, tree))
	assert.Equal("Root cannot have sibling", err.Error())

	parent := NewLinkedListTree(1)
	tree.SetParent(parent)
	err = tree.SetNextSibling(NewLinkedListTree(3))
	assert.Equal("Cannot set a root to be another tree's sibling", err.Error())

	sibling := NewLinkedListTreeNode(3, tree)
	err = tree.SetNextSibling(sibling)
	assert.Equal(sibling, tree.NextSibling())

	err = tree.SetNextSibling(nil)
	assert.Nil(tree.NextSibling())
}

func TestNextSibling(t *testing.T) {
	assert := assert.New(t)

	parent := NewLinkedListTree(1)
	tree := NewLinkedListTree(2)
	tree.SetParent(parent)
	sibling, _ := tree.AppendSibling(2)
	assert.Equal(sibling, tree.NextSibling())
}

func TestSetFirstChild(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(1)
	err := tree.SetFirstChild(NewLinkedListTree(2))
	assert.Equal("Cannot set a root to be another tree's child", err.Error())

	child := NewLinkedListTreeNode(2, tree)
	err = tree.SetFirstChild(child)
	assert.Equal(child, tree.FirstChild())

	err = tree.SetFirstChild(nil)
	assert.Nil(tree.FirstChild())
}

func TestNextChild(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(1)
	child := NewLinkedListTreeNode(2, tree)
	tree.SetFirstChild(child)
	assert.Equal(child, tree.FirstChild())
}

func ExamplePreOrderTraverse() {
	tree := NewLinkedListTree(1)
	child := tree.AppendChild(2)
	child.AppendChild(7)
	child.AppendChild(8)
	child = tree.AppendChild(3)
	child.AppendChild(9)
	child.AppendChild(10)
	child = tree.AppendChild(4)
	child.AppendChild(11)
	child.AppendChild(12)
	child.AppendChild(13)
	child = tree.AppendChild(5)
	child.AppendChild(14)
	child.AppendChild(15)
	child = tree.AppendChild(6)
	child.AppendChild(16)
	child.AppendChild(17)
	child.AppendChild(18)
	child.AppendChild(19)
	child.AppendChild(20)

	tree.Traverse(func(val interface{}, depth int) {
		for i, l := 0, depth*4; i < l; i++ {
			fmt.Print(" ")
		}
		fmt.Println(val)
	}, TRAVERSAL_PRE_ORDER)
	// Output:
	// 1
	//     2
	//         7
	//         8
	//     3
	//         9
	//         10
	//     4
	//         11
	//         12
	//         13
	//     5
	//         14
	//         15
	//     6
	//         16
	//         17
	//         18
	//         19
	//         20
}

func ExamplePostOrderTraverse() {
	tree := NewLinkedListTree(1)
	child := tree.AppendChild(2)
	child.AppendChild(7)
	child.AppendChild(8)
	child = tree.AppendChild(3)
	child.AppendChild(9)
	child.AppendChild(10)
	child = tree.AppendChild(4)
	child.AppendChild(11)
	child.AppendChild(12)
	child.AppendChild(13)
	child = tree.AppendChild(5)
	child.AppendChild(14)
	child.AppendChild(15)
	child = tree.AppendChild(6)
	child.AppendChild(16)
	child.AppendChild(17)
	child.AppendChild(18)
	child.AppendChild(19)
	child.AppendChild(20)

	tree.Traverse(func(val interface{}, depth int) {
		for i, l := 0, depth*4; i < l; i++ {
			fmt.Print(" ")
		}
		fmt.Println(val)
	}, TRAVERSAL_POST_ORDER)
	// Output:
	//         7
	//         8
	//     2
	//         9
	//         10
	//     3
	//         11
	//         12
	//         13
	//     4
	//         14
	//         15
	//     5
	//         16
	//         17
	//         18
	//         19
	//         20
	//     6
	// 1
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	tree := NewLinkedListTree(1)
	var clonedTree LinkedListTree
	clonedTree = *tree
	tree.Delete()
	assert.Equal(clonedTree, *tree)

	child := tree.AppendChild(2)
	child.Delete()
	assert.Nil(tree.FirstChild())

	tree.AppendChild(2)
	child = tree.AppendChild(3)
	child.Delete()
	assert.Nil(tree.FirstChild().NextSibling())
}
