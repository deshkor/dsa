package main

import (
	"fmt"
	"log"
)

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type List struct {
	size int
	head *Node
	tail *Node
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Add(val int) error {
	if l.Find(val) != nil {
		return fmt.Errorf("Element already exists.")
	}

	node := Node{Value: val}

	if l.head == nil {
		l.head = &node
	} else {
		l.tail.Next = &node
		node.Prev = l.tail
	}

	l.tail = &node
	l.size++

	return nil
}

func (l *List) Remove(val int) error {
	node := l.Find(val)

	if node == nil {
		return fmt.Errorf("Element not found.")
	}

	if node.Prev == nil {
		l.head = node.Next
	} else {
		node.Prev.Next = node.Next
		node.Prev = nil
	}

	if node.Next == nil {
		l.tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
		node.Next = nil
	}

	l.size--

	return nil
}

func (l *List) Find(val int) *Node {
	cur := l.head
	for cur != nil {
		if cur.Value == val {
			return cur
		}

		tmp := cur
		cur = tmp.Next
	}

	return nil
}

func (l *List) PrintAll() {
	log.Println("Size: ", l.size)

	cur := l.head
	for cur != nil {
		log.Println("Value:", cur.Value, "Prev:", cur.Prev, "Next:", cur.Next)

		tmp := cur
		cur = tmp.Next
	}
}
