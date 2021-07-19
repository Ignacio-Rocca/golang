package main

import (
	"fmt"
)

// a -> b -> c -> d

// a <- b <- c <- d
func reverseLinkedList(head *Node) *Node {
	if head == nil {
		return &Node{}
	}

	var arrValues []int
	tempNode := head
	for tempNode != nil {
		arrValues = append(arrValues, tempNode.Val)
		tempNode = tempNode.Next
	}

	fmt.Println(arrValues)
	return fromArrToReversedLinkedList(arrValues)
}

func fromArrToReversedLinkedList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	res := &Node{}
	for i := len(values)-1; i >= 0; i-- {
		appendNode(res, values[i])
	}

	return res
}

func appendNode(head *Node, val int) {
	if head.Next == nil {
		head.Val = val
		head.Next = &Node{}
		return
	}

	appendNode(head.Next, val)
}
