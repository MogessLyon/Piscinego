package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 1 {
		if f(a[0], a[1]) > 0 {
			for i := len(a) - 1; i > 0; i-- {
				if a[i-1] < a[i] {
					return false
				}
			}
		} else if f(a[0], a[1]) <= 0 {
			for i := 0; i < len(a)-1; i++ {
				if a[i] > a[i+1] {
					return false
				}
			}
		}
	}
	return true
}
