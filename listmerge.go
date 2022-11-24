package piscine

func ListMerge(l1 *List, l2 *List) {
	current := l1.Head
	str := current
	if current == nil {
		l1.Head = l2.Head
		return
	}
	for current != nil {
		str = current
		current = current.Next

	}
	str.Next = l2.Head
}
