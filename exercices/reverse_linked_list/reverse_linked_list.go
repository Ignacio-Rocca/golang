package main

import (
	"fmt"
)

type NodeList struct {
	Val int
	NextNode *NodeList
}

// a -> b -> c -> d

// a <- b <- c <- d
func reverseLinkedList(head *NodeList) *NodeList {
	if head == nil {
		return &NodeList{}
	}

	var arrValues []int
	tempNode := head
	for tempNode != nil {
		arrValues = append(arrValues, tempNode.Val)
		tempNode = tempNode.NextNode
	}

	fmt.Println(arrValues)
	return fromArrToReversedLinkedList(arrValues)
}

func fromArrToReversedLinkedList(values []int) *NodeList {
	if len(values) == 0 {
		return nil
	}
	res := &NodeList{}
	for i := len(values)-1; i >= 0; i-- {
		appendNode(res, values[i])
	}

	return res
}

func appendNode(head *NodeList, val int) {
	if head.NextNode == nil {
		head.Val = val
		head.NextNode = &NodeList{}
		return
	}

	appendNode(head.NextNode, val)
}

func main() {

	l := &NodeList{Val: 1, NextNode: &NodeList{Val: 2, NextNode: &NodeList{Val: 3}}}
	reversed := reverseLinkedList(l)

	for reversed != nil {
		fmt.Println(reversed.Val)
		reversed = reversed.NextNode
	}
}