package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root            *Node
	AllowDuplicates bool
	Size            int
}

func (t *Tree) Insert(value int) error {

	node := Node{Value: value}

	if t.Root == nil {
		t.Root = &node
		t.Size++

		return nil
	}

	return t.insertNode(t.Root, &node)
}

func (t *Tree) Search(value int) (*Node, error) {
	if t.Root == nil {
		return nil, errors.New("Cannot search an empty tree")
	}

	return t.searchTree(t.Root, value)
}

func (t *Tree) Remove(value int) error {
	return nil
}

func (t *Tree) insertNode(root, node *Node) error {
	if node == nil {
		return errors.New("Cannot insert a nil node.")
	}

	if root.Value > node.Value {
		if root.Left == nil {
			root.Left = node
		} else {
			return t.insertNode(root.Left, node)
		}
	}

	if root.Value <= node.Value {
		if root.Value == node.Value && !t.AllowDuplicates {
			return errors.New(fmt.Sprintf("Cannot insert duplicated value %d", node.Value))
		}

		if root.Right == nil {
			root.Right = node
		} else {
			return t.insertNode(root.Right, node)
		}
	}

	t.Size++

	return nil
}

func (t *Tree) searchTree(node *Node, value int) (*Node, error) {
	if node == nil {
		return nil, nil
	}

	if node.Value == value {
		return node, nil
	}

	if node.Value < value {
		return t.searchTree(node.Right, value)
	}

	return t.searchTree(node.Left, value)
}
