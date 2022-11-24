package piscine

func Compare(a, b string) int {
	var comp int
	if a > b {
		comp = 1
	} else if a < b {
		comp = -1
	} else {
		comp = 0
	}
	return comp
}
