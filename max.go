package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	SortIntegerTable(a)
	return a[len(a)-1]
}
