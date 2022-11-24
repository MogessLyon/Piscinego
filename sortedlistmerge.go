package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	current := n1
	str := n1
	if current == nil {
		str = n2
		return nil
	}

	for current != nil {
		str = current
		current = current.Next
	}

	str.Next = n2

	return ListSort(n1)
}
