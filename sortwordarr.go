package piscine

func SortIntegerByBytes(table []string) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			tmp := table[i-1]
			tmp1 := table[i]
			table[i-1] = tmp1
			table[i] = tmp
			i = 1
		} else {
			i++
		}
	}
}

func SortWordArr(a []string) {
	SortIntegerByBytes(a)
}
