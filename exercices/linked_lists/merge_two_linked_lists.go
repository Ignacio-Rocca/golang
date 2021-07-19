package main



func mergeTwoLists(l1 *Node, l2 *Node) *Node {
	if l1 == nil && l2 == nil { return &Node{} }

	if l1 == nil { return l2 }

	if l2 == nil { return l1 }

	if l1.Val == 0 && l2.Val == 0 {
		return &Node{}
	}

	result := &Node{}
	if l1.Val <= l2.Val {
		result.Val = l1.Val
		l1 = l1.Next
	} else if l1.Val > l2.Val {
		result.Val = l2.Val
		l2 = l2.Next
	}

	lastNode := &Node{}
	result.Next = lastNode

	for l1.Next != nil || l2.Next != nil  {
		if l1.Val <= l2.Val {
			lastNode.Val = l1.Val
			l1 = l1.Next
		} else {
			lastNode.Val = l2.Val
			l2 = l2.Next
		}

		lastNode.Next = &Node{}
		lastNode = lastNode.Next
	}

	if l1.Val <= l2.Val {
		lastNode.Val = l1.Val
		lastNode.Next = &Node{Val: l2.Val}
	} else {
		lastNode.Val = l2.Val
		lastNode.Next = &Node{Val: l1.Val}
	}

	return result
}
