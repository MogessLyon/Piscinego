package piscine

func ListLast(l *List) interface{} {
	p := l.Head
	if p == nil {
		return nil
	}
	for p.Next != nil {
		p = p.Next
		if p.Next == nil {
			return p.Data
		}
	}

	return nil
}
