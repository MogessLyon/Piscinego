package piscine

func ListSize(l *List) int {
	p := l.Head
	ln := 0
	for p != nil {
		p = p.Next
		ln++
	}

	return ln
}
