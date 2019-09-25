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
	node, err := t.Search(value)
	if node == nil || err != nil {
		return fmt.Errorf("Can't delete %d from tree.", value)
	}

	parent, err := t.findParent(node)
	if err != nil {
		return err
	}

	// node is the root TODO: can this be reused?
	if parent == nil {
		if node.Left == nil {
			t.Root = node.Right
		} else if node.Right == nil {
			t.Root = node.Left
		} else {
			leastRightNode := node.Right
			for leastRightNode.Left != nil {
				cur := leastRightNode
				leastRightNode = cur.Left
			}

			if leastRightNode == nil {
				return fmt.Errorf("Invalid value while trying to delete.")
			}

			leastRightNode.Left = node.Left
			t.Root = node.Right
		}

		t.Size--
		return nil
	}

	// node to remove is on the left
	if parent.Left == node {
		return nil
	}

	return nil
}

func (t *Tree) findParent(node *Node) (*Node, error) {
	// the node is the root
	if t.Root.Value == node.Value {
		return nil, nil
	}

	// checking if the root is the parent
	if t.Root.Left.Value == node.Value || t.Root.Right.Value == node.Value {
		return t.Root, nil
	}

	var current *Node

	if t.Root.Value < node.Value {
		current = t.Root.Right
	} else {
		current = t.Root.Left
	}

	if current == nil {
		return nil, fmt.Errorf("Unexpected value when setting up the tree to be searched.")
	}

	for {
		// the current value is the parent
		if current.Left.Value == node.Value || current.Right.Value == node.Value {
			break
		}

		tmp := current
		if node.Value > current.Value {
			current = tmp.Right
		} else {
			current = tmp.Left
		}

		if current == nil {
			return nil, fmt.Errorf("Unexpected when searching for parent.")
		}
	}

	return current, nil
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
