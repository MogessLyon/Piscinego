package piscine

func ListPushFront(l *List, data interface{}) {
	list := &NodeL{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = list, l.Head
	} else {
		list.Next, l.Head = l.Head, list
	}
}
