package main

func subListSearch(headList *Node, headSubList *Node) bool {
	if headList == nil && headSubList == nil {
		return true
	}

	if headList == nil || headSubList == nil {
		return false
	}

	if headList.Val == headSubList.Val {
		if headSubList.Next == nil {
			return true
		} else if headList.Next == nil {
			return false
		}
		return subListSearch(headList.Next, headSubList.Next)
	}

	return subListSearch(headList.Next, headSubList)
}
