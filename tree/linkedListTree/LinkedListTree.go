package linkedListTree

import (
	"errors"

	. "github.com/yuhlau/go-data-structures/tree"
)

type LinkedListTree struct {
	val         interface{}
	parent      *LinkedListTree
	prevSibling *LinkedListTree
	nextSibling *LinkedListTree
	firstChild  *LinkedListTree
}

// NewLinkedListTree creates and returns the pointer to the Linked List Tree
// with the specified value. The tree created is default to be a root
func NewLinkedListTree(val interface{}) *LinkedListTree {
	return &LinkedListTree{val: val}
}

// NewLinkedListTreeNode creates and returns the pointer to the Linked List
// Tree node with the specified value
func NewLinkedListTreeNode(val interface{}, parent *LinkedListTree) *LinkedListTree {
	return &LinkedListTree{parent: parent, val: val}
}

// SetVal updates the Linked List Tree value
func (tree *LinkedListTree) SetVal(val interface{}) {
	tree.val = val
}

// Val returns the value of the Linked List Tree
func (tree *LinkedListTree) Val() interface{} {
	return tree.val
}

// SetParent updates the pointer to the parent of the tree, returns an error
// if the node already has siblings
func (tree *LinkedListTree) SetParent(parent *LinkedListTree) error {
	if parent == nil && tree.NextSibling() != nil {
		return errors.New("Tree cannot be root because it already has sibling")
	}
	tree.parent = parent
	return nil
}

// Parent returns the pointer to the parent of the tree
func (tree *LinkedListTree) Parent() *LinkedListTree {
	return tree.parent
}

// IsRoot returns whether the tree is the root
func (tree *LinkedListTree) IsRoot() bool {
	return tree.Parent() == nil
}

// AppendSibling appends a sibilng to the current tree and returns the pointer
// to the created sibling, error if the tree is the root
func (tree *LinkedListTree) AppendSibling(val interface{}) (*LinkedListTree, error) {
	if tree.IsRoot() {
		return nil, errors.New("Root cannot have sibling")
	}
	// The sibling will have the same parent as this tree
	sibling := NewLinkedListTreeNode(val, tree.Parent())
	// append the sibling to the end of nextSibling list
	current := tree
	for current.NextSibling() != nil {
		current = current.NextSibling()
	}
	sibling.SetPrevSibling(current)
	err := current.SetNextSibling(sibling)
	if err != nil {
		return nil, err
	}

	return sibling, nil
}

// AppendChild appends a child to the current tree and returns the pointer to
// the created child
func (tree *LinkedListTree) AppendChild(val interface{}) *LinkedListTree {
	// Child will have the current tree as parent
	child := NewLinkedListTreeNode(val, tree)

	if tree.FirstChild() == nil {
		// This is the first child of the tree
		tree.firstChild = child
		return child
	}

	current := tree.FirstChild()
	for current.NextSibling() != nil {
		current = current.NextSibling()
	}
	child.SetPrevSibling(current)
	current.SetNextSibling(child)
	return child
}

// SetPrevSibling updates the pointer to the previous sibling of the tree with
// the provided Linked List Tree node, returns error if the tree is a root, or
// the provided sibling is a root
func (tree *LinkedListTree) SetPrevSibling(prevSibling *LinkedListTree) error {
	if prevSibling != nil {
		if tree.IsRoot() {
			return errors.New("Root cannot have sibling")
		}
		if prevSibling.IsRoot() {
			return errors.New("Cannot set a root to be another tree's sibling")
		}
	}
	tree.prevSibling = prevSibling
	return nil
}

// PrevSibling returns the pointer to the previous sibling of the tree
func (tree *LinkedListTree) PrevSibling() *LinkedListTree {
	return tree.prevSibling
}

// SetNextSibling updates the pointer to the next sibling of the tree with the
// provided Linked List Tree node, returns error if the tree is a root, or the
// provided sibling is a root
func (tree *LinkedListTree) SetNextSibling(nextSibling *LinkedListTree) error {
	if nextSibling != nil {
		if tree.IsRoot() {
			return errors.New("Root cannot have sibling")
		}
		if nextSibling.IsRoot() {
			return errors.New("Cannot set a root to be another tree's sibling")
		}
	}
	tree.nextSibling = nextSibling
	return nil
}

// NextSibling returns the pointer to the next sibling of the tree
func (tree *LinkedListTree) NextSibling() *LinkedListTree {
	return tree.nextSibling
}

// SetFirstChild updates the pointer to the first child node, return errors if
// the child is a root
func (tree *LinkedListTree) SetFirstChild(firstChild *LinkedListTree) error {
	if firstChild != nil {
		if firstChild.IsRoot() {
			return errors.New("Cannot set a root to be another tree's child")
		}
	}
	tree.firstChild = firstChild
	return nil
}

// FirstChild returns the pointer to the first child
func (tree *LinkedListTree) FirstChild() *LinkedListTree {
	return tree.firstChild
}

func (tree *LinkedListTree) _preOrderTraverse(fn func(interface{}, int), method, depth int) {
	fn(tree.Val(), depth)
	if child := tree.FirstChild(); child != nil {
		child._preOrderTraverse(fn, method, depth+1)
	}
	if sibling := tree.NextSibling(); sibling != nil {
		sibling._preOrderTraverse(fn, method, depth)
	}
}

func (tree *LinkedListTree) _postOrderTraverse(fn func(interface{}, int), method, depth int) {
	if child := tree.FirstChild(); child != nil {
		child._postOrderTraverse(fn, method, depth+1)
	}
	fn(tree.Val(), depth)
	if sibling := tree.NextSibling(); sibling != nil {
		sibling._postOrderTraverse(fn, method, depth)
	}
}

func (tree *LinkedListTree) Traverse(fn func(interface{}, int), method int) error {
	switch method {
	case TRAVERSAL_PRE_ORDER:
		tree._preOrderTraverse(fn, method, 0)
	case TRAVERSAL_POST_ORDER:
		tree._postOrderTraverse(fn, method, 0)
	default:
		return errors.New("Unsupported traversal method")
	}
	return nil
}

// Delete deletes the current node from the tree structure. If this is a tree
// node then it is removed from its parent and siblings; If this is a root
// then nothing would happen. You should always remove the reference to the
// tree to allow for garbage collection
func (tree *LinkedListTree) Delete() {
	if tree.IsRoot() {
		return
	}
	// Remove itself from siblings list and/or parent
	if tree.PrevSibling() == nil {
		// The tree is the first child in its parent
		tree.Parent().SetFirstChild(tree.NextSibling())
	} else {
		tree.PrevSibling().SetNextSibling(tree.NextSibling())
	}
}
