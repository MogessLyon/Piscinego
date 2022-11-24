package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head
	prev := current

	for current != nil {
		next := current.Next
		if current.Data == data_ref {
			if prev.Data == data_ref {
				l.Head = current.Next
			} else {
				prev.Next = next
			}
		} else {
			prev = current
		}
		current = next
	}
}
