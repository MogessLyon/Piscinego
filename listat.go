package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	i := 0
	for l != nil {

		if pos == i {
			return l
		}
		l = l.Next
		i++
	}

	return nil
}
