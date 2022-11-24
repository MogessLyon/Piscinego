package piscine

func Power(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	recursive := nb * Power(nb, power-1)
	return recursive
}

func Sqrt(nb int) int {
	a := 0
	if nb < 0 {
		return 0
	}
	for a < nb {
		if nb == Power(a, 2) {
			return a
		} else {
			a++
		}
	}
	if nb != 1 && nb == a {
		return 0
	}
	return a
}
