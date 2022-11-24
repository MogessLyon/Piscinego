package piscine

func ActiveBits(n int) int {
	var counter int
	for n != 0 {
		if n%2 == 1 {
			counter++
		}
		n /= 2
	}
	return counter
}
