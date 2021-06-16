package merge_two_linked_lists

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil { return &ListNode{} }

	if l1 == nil { return l2 }

	if l2 == nil { return l1 }

	if l1.Val == 0 && l2.Val == 0 {
		return &ListNode{}
	}

	result := &ListNode{}
	if l1.Val <= l2.Val {
		result.Val = l1.Val
		l1 = l1.Next
	} else if l1.Val > l2.Val {
		result.Val = l2.Val
		l2 = l2.Next
	}

	lastNode := &ListNode{}
	result.Next = lastNode

	for l1.Next != nil || l2.Next != nil  {
		if l1.Val <= l2.Val {
			lastNode.Val = l1.Val
			l1 = l1.Next
		} else {
			lastNode.Val = l2.Val
			l2 = l2.Next
		}

		lastNode.Next = &ListNode{}
		lastNode = lastNode.Next
	}

	if l1.Val <= l2.Val {
		lastNode.Val = l1.Val
		lastNode.Next = &ListNode{Val: l2.Val}
	} else {
		lastNode.Val = l2.Val
		lastNode.Next = &ListNode{Val: l1.Val}
	}

	return result
}
