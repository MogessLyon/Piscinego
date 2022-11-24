package piscine

func Map(f func(int) bool, a []int) []bool {
	var arr []bool
	for _, value := range a {
		arr = append(arr, f(value))
	}
	return arr
}
