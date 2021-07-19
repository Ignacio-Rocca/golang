package main

import (
	"fmt"
)

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

type Node struct {
	Val int
	Next *Node
}

// Push insert at front a new node
func (n *Node) Push(val int) *Node{
	return &Node{Val: val, Next: n}
}

func (l LinkedList) Len() int {
	return l.length
}

func (l LinkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.Val)
		l.head = l.head.Next
	}
	fmt.Println()
}

func (l LinkedList) Front() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot Find Front Value in an Empty linked list")
	}
	return l.head.Val, nil
}

func (l LinkedList) Back() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot Find Front Value in an Empty linked list")
	}
	return l.tail.Val, nil
}

func (l *LinkedList) Reverse() {
	curr := l.head
	l.tail = l.head
	var prev *Node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *LinkedList) PushBack(n *Node) {
	if l.head == nil {
		l.head = n
	} else {
		l.tail.Next = n
	}
	l.tail = n
	l.length++
}

func (l *LinkedList) Delete(key int) {
	if l.head.Val == key {
		l.head = l.head.Next
		l.length--
		return
	}

	var prev *Node = nil
	curr := l.head
	for curr != nil && curr.Val != key {
		prev = curr
		curr = curr.Next
	}

	if curr == nil {
		fmt.Println("Key Not found")
		return
	}

	prev.Next = curr.Next
	l.length--
	fmt.Println("Node Deleted")
}

func CopyNodesList(head *Node) *Node {
	var newList *Node
	var tail *Node
	current := head

	for current != nil {
		if newList == nil {
			newList = current
			tail = newList
		} else {
			tail = tail.Next
			tail.Val = current.Val
			tail.Next = nil
		}
		current = current.Next
	}

	return newList
}