package piscine

func Unmatch(a []int) int {
	SortIntegerTable(a)
	value := -1

	for idx, val := range a {
		counter := 1
		for i, v := range a {
			if idx != i && val == v {
				counter++
			}
		}
		if counter%2 != 0 {
			value = val
			break
		}
		counter = 0
	}
	return value
}
