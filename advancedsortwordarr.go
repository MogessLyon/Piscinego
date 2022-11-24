package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	i := 1
	for i < len(a) {
		v := string(a[i-1])
		k := string(a[i])

		if f(v, k) == 1 {
			tmp := a[i-1]
			tmp1 := a[i]
			a[i-1] = tmp1
			a[i] = tmp
			i = 1
		} else {
			i++
		}
	}
}
