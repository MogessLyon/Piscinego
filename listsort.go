package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l
	if l == nil {
		return nil
	}

	for current.Next != nil {
		if current.Data > current.Next.Data {
			tmp := current.Data
			current.Data = current.Next.Data
			current.Next.Data = tmp
			current = l
		} else {
			current = current.Next
		}
	}
	return l
}
